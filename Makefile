BIN=server
BIN_PATH=build/$(BIN)
GOBIN=$(GOPATH)/bin

.PHONY: setup-tools
setup-tools: setup-reflex setup-goimports setup-migrate

.PHONY: setup-reflex
setup-reflex: $(GOBIN)/reflex

$(GOBIN)/reflex:
	cd && go get github.com/cespare/reflex

.PHONY: setup-goimports
setup-goimports: $(GOBIN)/goimports

$(GOBIN)/goimports:
	cd && go get golang.org/x/tools/cmd/goimports

.PHONY: setup
setup:
	go mod download

.PHONY: tidy-modules
tidy-modules:
	go mod tidy -v

.PHONY: build
build: setup
	go build -o $(BIN_PATH) -a -tags netgo -installsuffix netgo --ldflags '-extldflags "-static"' ./

.PHONY: run
run: build
	$(BIN_PATH)

.PHONY: localup
localup: setup-reflex
	$(GOBIN)/reflex -r '(\.go|go\.mod)' -d none -s make run

.PHONY: clean
clean:
	rm -rf build

.PHONY: goimports
goimports: setup-goimports
	$(GOBIN)/goimports -w .

.PHONY: gqlgen
gqlgen:
	go run github.com/99designs/gqlgen -v
