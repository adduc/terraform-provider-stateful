build:
	go build -o bin/$(shell basename $(PWD)) cmd/main.go

testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

fmt:
	go fmt ./...

plan:
	cd examples && terraform plan

apply:
	cd examples && terraform apply