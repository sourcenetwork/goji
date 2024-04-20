.PHONY: deps\:test
deps\:test:
	go install github.com/agnivade/wasmbrowsertest@latest

.PHONY: test\:js
test:
	GOOS=js GOARCH=wasm go test -timeout 30s -exec wasmbrowsertest ./...