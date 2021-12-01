# Set the shell to bash always
SHELL := /bin/bash

init=@./make-tools --init -m "starting"
save=@./make-tools --message "saving value" --save
load=@./make-tools --debug --load
exists=@./make-tools -m "environment variable doesn't exist" --env-exists
X=load -l hello
# Options

export ENV_SECRET="secret"
t:
	go build -o make-tools main.go
	$(init)
	$(exists) ENV_SECRET
	$(save) encoded1=`echo -n yyy | base64`
	$(save) encoded2=`echo bubu | base64`
	$(load) encoded1
	@echo
	$(save) hello=SGVsbG8gZnJvbSBFTkNPREVE
	@echo "the encoded message: `./make-tools -l hello | base64 -d`"



build: test
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o ./bin/make-tools main.go

check: lint test

lint:
	$(LINT) run

tidy:
	go mod tidy

test:
	go test -v ./...

# Tools

LINT=$(shell which golangci-lint)

.PHONY: tidy lint build run
