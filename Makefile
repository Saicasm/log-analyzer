# Variables
GOCMD = go
GOTEST = $(GOCMD) test
GOBUILD = $(GOCMD) build
GOINSTALL = $(GOCMD) install
GOCOVER = $(GOCMD) tool cover
GOCLEAN = $(GOCMD) clean
DEPCMD = dep
CLI_BINARY_NAME = cli

# Targets
all: install build

install:
	$(DEPCMD) ensure

build: cli server

cli:
	$(GOBUILD) -o $(CLI_BINARY_NAME) cmd/main.go cmd/cli.go cmd/server.go


run_cli: cli
	@echo "Running CLI..."
	@export LOG_MODE=CLI && ./$(CLI_BINARY_NAME) $(ARGS)
#  make run_cli ARGS="-f cookie_log.csv -d 2018-12-09 "
run_server: cli
	@echo "Running Server..."
	@export LOG_MODE=SERVER && ./$(CLI_BINARY_NAME) $(ARGS)

test:
	$(GOTEST) ./... -coverprofile=coverage.out

coverage: test
	$(GOCOVER) -html=coverage.out

clean:
	$(GOCLEAN)
	rm -f $(CLI_BINARY_NAME) $(SERVER_BINARY_NAME) coverage.out
