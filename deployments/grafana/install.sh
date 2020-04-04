#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

VALUES="${DIR}/values.yaml"

kubectl create namespace grafana
helm upgrade --install -n grafana -f ${VALUES} grafana k8sgoo/grafana