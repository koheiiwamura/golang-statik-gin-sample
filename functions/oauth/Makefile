GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_DIR=../../bin/


build:
	statik -src=templates/
	$(GOBUILD) -o $(BINARY_DIR) .

build-linux:
	statik -src=templates/
	GOOS=linux $GOARCH=386 $(GOBUILD) -o $(BINARY_DIR) .
