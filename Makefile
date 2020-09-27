.PHONY: list
list:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'
ecr-login:
	@aws ecr get-login-password --region eu-west-1 | docker login --username AWS --password-stdin 888435310358.dkr.ecr.eu-west-1.amazonaws.com
ecr-create-repo: ecr-login
	@aws ecr create-repository --repository-name windstats/github.com/costap/windstats/cmd/windstatsapi
	@aws ecr create-repository --repository-name windstats/github.com/costap/windstats/cmd/windstatsd
apply: ecr-login
	@export KO_DOCKER_REPO=888435310358.dkr.ecr.eu-west-1.amazonaws.com/windstats; kustomize build deployments/ko/local-kubepi | ko apply -P -f -

generate:
	@protoc internal/repository/repository.proto --go_out=plugins=grpc:.

pf-db:
	@kubectl port-forward -n windstats deploy/influxdb 8086

run:
	@export WINDSTATS_DBADDR=http://localhost:8086; go run cmd/windstatsd/* cmd/windstatsd/app/*

ecr-secret:
	@ACCOUNT=888435310358 REGION=eu-west-1 SECRET_NAME=ecr NS=windstats ./ecr-secret.sh
