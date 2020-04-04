.PHONY: list
list:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'

start-registry:
	@docker run -d -p 5000:5000 --restart=always --name registry registry:2

apply:
	@export KO_DOCKER_REPO=registry:5000; kustomize build deployments/ko/local-kubepi | ko apply --insecure-registry -f -

generate:
	@protoc internal/repository/repository.proto --go_out=plugins=grpc:.

pf-db:
	@kubectl port-forward -n windstats deploy/influxdb 8086

run:
	@export WINDSTATS_DBADDR=http://localhost:8086; go run cmd/windstatsd/* cmd/windstatsd/app/*
