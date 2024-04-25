#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname ${BASH_SOURCE[0]})/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd ${SCRIPT_ROOT}; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}

bash ${CODEGEN_PKG}/generate-internal-groups.sh all \
  github.com/setcreed/store-apiserver/pkg/generated github.com/setcreed/store-apiserver/pkg/apis github.com/setcreed/store-apiserver/pkg/apis \
  "store:v1alpha1" \
  --go-header-file ${SCRIPT_ROOT}/hack/boilerplate.go.txt
