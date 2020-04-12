.PHONY: list
list:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'

build-init:
	@docker buildx use testbuilder
	@docker buildx inspect --bootstrap

build:
	@docker buildx build --platform linux/amd64,linux/arm/v7 -f deployments/docker/Dockerfile --output type=image,name=host.docker.internal:5000/windstats-web:latest,push=true,registry.insecure=true .

run:
	@docker run --name windstatsweb -d -p 8080:80 registry:5000/windstats-web:latest

create-builder:
	@docker buildx create --name testbuilder
