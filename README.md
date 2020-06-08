## Application with NATS streaming for learning and exploration  ##

This application is based on the works of Tin Rabzelj

Building a Microservices Application in Go Following the CQRS Pattern
https://outcrawl.com/go-microservices-cqrs-docker
https://github.com/tinrab/meower


Technologies involved:

NATS streaming
Elasticsearch
Postgres
nginx
docker-compose
vuejs
websocker


To run:

1. $ make up

2. $ make push ( then change item in Makefile push: multiple times)

3. $ make query-meows (to see all the items)

4. $ make query-search (to search particular items)

5. $ make tail-meow (to see logs)

6. $ make tail-pusher (to see logs)

7. $ make tail-query (to see logs)