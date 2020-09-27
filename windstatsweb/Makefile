.PHONY: list
list:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'

build-init:
	@docker buildx use testbuilder
	@docker buildx inspect --bootstrap

build:
	@docker buildx build --platform linux/arm/v7 -f deployments/docker/Dockerfile --load -t 888435310358.dkr.ecr.eu-west-1.amazonaws.com/windstats/windstatsweb:latest .
	@docker push 888435310358.dkr.ecr.eu-west-1.amazonaws.com/windstats/windstatsweb:latest

run:
	@docker run --name windstatsweb -d -p 8080:80 registry:5000/windstats-web:latest

create-builder:
	@docker buildx create --name testbuilder

ecr-login:
	@aws ecr get-login-password --region eu-west-1 | docker login --username AWS --password-stdin 888435310358.dkr.ecr.eu-west-1.amazonaws.com
ecr-create-repo: ecr-login
	@aws ecr create-repository --repository-name windstats/windstatsweb
