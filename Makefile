SRC_DIR := "/go/src/kubectl-switch"
OS ?= "darwin"
ARCH ?= "amd64"

build:
	go build -o bin/kubectl-switch

build-docker:
	docker run -v $$(pwd):$(SRC_DIR) -w /go/src/kubectl-switch \
		-e GOOS=$(OS) -e GOARCH=$(ARCH) golang:1.14.1-alpine \
		go build -o $(SRC_DIR)/bin/kubectl-switch
