up:
	docker-compose up --build -d
.PHONY: up


tail-meow:
	docker logs meow-service -f 
tail-pusher:
	docker logs pusher-service -f 
tail-query:
	docker logs query-service -f 
.PHONY: tail-meow tail-pusher tail-query 

prune:
	docker system prune -a --volumes
.PHONY: prune

query-meows:
	curl http://localhost:7474/meows 
query-search:
	curl http://localhost:7474/search?query=rambutans&skip=0&take=5
.PHONY: query-meows query-search

push:
	curl -X POST http://localhost:8080/meows?body=rambutans
.PHONY: push