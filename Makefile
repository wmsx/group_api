.PHONY: docker
docker:
	docker build . -t group-api:latest
