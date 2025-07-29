.PHONY: deps\:test
deps\:test:
	go install github.com/agnivade/wasmbrowsertest@latest
	go install github.com/sourcenetwork/web-transport-echo-server@latest

.PHONY: test
test:
	./run-tests.sh
