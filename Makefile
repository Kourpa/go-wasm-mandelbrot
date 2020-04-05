GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=mandelbrot.wasm
MAIN=main.go
RUNNER=server.go

all: build-wasm
build-wasm:
	GOOS=js GOARCH=wasm $(GOBUILD) -o $(BINARY_NAME) $(MAIN)
	mv $(BINARY_NAME) ./web/
run: build-wasm
	$(GORUN) $(RUNNER)
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f ./web/$(BINARY_NAME)
