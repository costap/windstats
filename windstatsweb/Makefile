.PHONY: list
list:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'

build-init:
	@docker buildx use testbuilder
	@docker buildx inspect --bootstrap

build:
	@docker buildx build --platform linux/arm/v7 -f deployments/docker/Dockerfile --load -t pedrofcosta/windstatsweb:latest .
	@docker push pedrofcosta/windstatsweb:latest

create-builder:
	@docker buildx create --name testbuilder
