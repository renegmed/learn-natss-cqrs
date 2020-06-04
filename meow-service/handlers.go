/*
  microservices cqrs pattern tin rabzelj
*/
package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/renegmed/microserv-cqrs-natss/meow-service/db"
	"github.com/renegmed/microserv-cqrs-natss/meow-service/event"
	"github.com/renegmed/microserv-cqrs-natss/meow-service/schema"
	"github.com/renegmed/microserv-cqrs-natss/meow-service/util"
	"github.com/segmentio/ksuid"
)

func createMeowHandler(w http.ResponseWriter, r *http.Request) {
	type response struct {
		ID string `json:"id"`
	}

	ctx := r.Context()

	// Read parameters e.g. 'body' not empty and length not more than 140 characters
	body := template.HTMLEscapeString(r.FormValue("body"))
	if len(body) < 1 || len(body) > 140 {
		util.ResponseError(w, http.StatusBadRequest, "Invalid body")
		return
	}

	log.Printf("createMeowHandler body:\n %v\n", body)

	// Create meow
	createdAt := time.Now().UTC()
	// Generate ID
	id, err := ksuid.NewRandomWithTime(createdAt)
	if err != nil {
		util.ResponseError(w, http.StatusInternalServerError, "Failed to create meow")
		return
	}
	meow := schema.Meow{
		ID:        id.String(),
		Body:      body,
		CreatedAt: createdAt,
	}

	// add record to Postgres db
	if err := db.InsertMeow(ctx, meow); err != nil {
		log.Println(err)
		util.ResponseError(w, http.StatusInternalServerError, "Failed to create meow")
		return
	}

	// Publish nats event to notify others
	if err := event.PublishMeowCreated(meow); err != nil {
		log.Println(err)
	}

	// Return new meow
	util.ResponseOk(w, response{ID: meow.ID})
}
