#!/bin/bash

trap "kill 0" EXIT

web-transport-echo-server &

GOOS=js GOARCH=wasm go test -cover -timeout 30s -exec wasmbrowsertest ./...
