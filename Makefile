GO15VENDOREXPERIMENT=1

NAME	 := pexpo
TARGET	 := bin/$(NAME)
VERSION  := 1.41
DIST_DIRS := find * -type d -exec

SRCS	:= $(shell find . -type f -name '*.go')
LDFLAGS := -ldflags="-s -w -X \"main.version=$(VERSION)\" -extldflags \"-static\""

$(TARGET): $(SRCS)
	go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o bin/$(NAME)

.PHONY: install
install:
	go install $(LDFLAGS)

.PHONY: wash
wash:
	rm -rf bin/*
	rm -rf dist/*

.PHONY: clean
clean:
	rm -rf bin/*
	rm -rf vendor/*
	rm -rf dist/*

.PHONY: cross-build
cross-build: deps
	for os in darwin linux windows; do \
		for arch in amd64 386; do \
			GOOS=$$os GOARCH=$$arch CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o dist/$(NAME)-$$os-$$arch/$(NAME); \
		done; \
	done

.PHONY: upde
upde:
	glide up

.PHONY: deps
deps:
	glide install

.PHONY: glide
glide:
ifeq ($(shell command -v glide 2> /dev/null),)
	go get github.com/Masterminds/glide
	go install github.com/Masterminds/glide
endif

.PHONY: dist
dist:
	cd dist && \
	$(DIST_DIRS) cp ../LICENSE {} \; && \
	$(DIST_DIRS) cp ../README.md {} \; && \
	$(DIST_DIRS) cp ../ping-list.txt {} \; && \
	$(DIST_DIRS) cp ../curl-list.txt {} \; && \
	$(DIST_DIRS) tar -zcf $(NAME)-$(VERSION)-{}.tar.gz {} \; && \
	$(DIST_DIRS) zip -r $(NAME)-$(VERSION)-{}.zip {} \; && \
	cd ..
