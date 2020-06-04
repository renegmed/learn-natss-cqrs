/*
  microservices cqrs pattern tin rabzelj
*/
package search

import (
	"context"
	"encoding/json"
	"log"

	"github.com/renegmed/microserv-cqrs-natss/query-service/schema"

	"github.com/olivere/elastic"
)

type ElasticRepository struct {
	client *elastic.Client
}

func NewElastic(url string) (*ElasticRepository, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}
	return &ElasticRepository{client}, nil
}

func (r *ElasticRepository) Close() {
}

func (r *ElasticRepository) InsertMeow(ctx context.Context, meow schema.Meow) error {
	_, err := r.client.Index().
		Index("meows").
		Type("meow").
		Id(meow.ID).
		BodyJson(meow).
		Refresh("wait_for").
		Do(ctx)
	return err
}

func (r *ElasticRepository) SearchMeows(ctx context.Context, query string, skip uint64, take uint64) ([]schema.Meow, error) {
	result, err := r.client.Search().
		Index("meows").
		Query(
			elastic.NewMultiMatchQuery(query, "body").
				Fuzziness("3").
				PrefixLength(1).
				CutoffFrequency(0.0001),
		).
		From(int(skip)).
		Size(int(take)).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	meows := []schema.Meow{}
	for _, hit := range result.Hits.Hits {
		var meow schema.Meow
		/*

			type Meow struct {
				ID        string    `json:"id"`
				Body      string    `json:"body"`
				CreatedAt time.Time `json:"created_at"`
			}

		*/
		if err = json.Unmarshal(*hit.Source, &meow); err != nil {
			log.Println(err)
		}
		meows = append(meows, meow)
	}
	return meows, nil
}
