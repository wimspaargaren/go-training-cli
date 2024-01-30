.PHONY: run lint gofumpt gci test postgres-start postgres-stop

run:
	go run ./cmd/cli

.PHONY: lint

$(GOBIN)/golangci-lint:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2
	@go mod tidy

lint: | $(GOBIN)/golangci-lint
	@golangci-lint  -v --concurrency=3 --config=.golangci.yml --issues-exit-code=1 run \
	--out-format=colored-line-number

$(GOBIN)/gofumpt:
	@go install mvdan.cc/gofumpt@v0.5.0
	@go mod tidy

gofumpt: | $(GOBIN)/gofumpt
	@gofumpt -w $(shell ls  -d $(PWD)/*/)

$(GOBIN)/gci:
	@go install github.com/daixiang0/gci@v0.11.0
	@go mod tidy

gci: | $(GOBIN)/gci
	@gci write --section Standard --section Default --section "Prefix(github.com/wimspaargaren/go-training-cli)" $(shell ls  -d $(PWD)/*)

# Run unit tests and generate coverage report
test:
	@mkdir -p reports
	@go test -coverprofile=reports/codecoverage_all.cov ./... -cover -race -p=4
	@go tool cover -func=reports/codecoverage_all.cov > reports/functioncoverage.out
	@go tool cover -html=reports/codecoverage_all.cov -o reports/coverage.html
	@echo "View report at $(PWD)/reports/coverage.html"
	@tail -n 1 reports/functioncoverage.out

postgres-start:
	@docker-compose -f ./docker-compose.yaml up -d

postgres-stop:
	@docker-compose -f ./docker-compose.yaml down