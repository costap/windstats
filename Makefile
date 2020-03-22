.PHONY: list
list:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'

apply:
	@export KO_DOCKER_REPO=registry:5000; kustomize build deployments/ko/local-kubepi | ko apply --insecure-registry -f -

generate:
	@protoc internal/repository/repository.proto --go_out=plugins=grpc:.

pf-db:
	@kubectl port-forward -n windstats svc/influxdb 8086
