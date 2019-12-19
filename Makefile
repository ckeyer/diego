PWD := $(shell pwd)
APP := diego
PKG := github.com/ckeyer/$(APP)

GO := go
HASH := $(shell which shasum || which sha1sum)

OS := $(shell go env GOOS)
ARCH := $(shell go env GOARCH)
VERSION := $(shell cat VERSION.txt)
GIT_COMMIT := $(shell git rev-parse --short HEAD)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
BUILD_AT := $(shell date "+%Y-%m-%dT%H:%M:%SZ%z")
PACKAGE_NAME := $(APP)$(VERSION).$(OS)-$(ARCH)

COMMONS_PKG := $(PKG)/vendor/github.com/ckeyer/commons
LD_FLAGS := -X ${COMMONS_PKG}/version.version=$(VERSION) \
 -X ${COMMONS_PKG}/version.gitCommit=$(GIT_COMMIT) \
 -X ${COMMONS_PKG}/version.buildAt=$(BUILD_AT)

IMAGE := ckeyer/${APP}
GO_IMAGE := ckeyer/dev:go
UI_IMAGE := ckeyer/dev:vue

env:
	$(GO) env

gorun:
	$(GO) run -ldflags="$(LD_FLAGS)" main.go

build: env generate-ui
	$(GO) build -v -ldflags="$(LD_FLAGS)" -o bundles/$(APP) main.go
	$(HASH) bundles/$(APP)

test:
	$(GO) test -cover -covermode=count $$(go list ./... |grep -v "vendor")

image:
	docker build -t ${IMAGE}:${VERSION} .

push-image: image
	docker push ${IMAGE}:${VERSION}
	@echo "push image ${IMAGE}:${VERSION} successful."

test-in-docker:
	docker run --rm \
	 -v ${PWD}:/go/src/${PKG} \
	 -w /go/src/${PKG} \
	 ${GO_IMAGE} \
	 go test -ldflags="$(LD_FLAGS)" $$(go list ./... |grep -v "vendor")

generate-ui:
	which go-bindata || go get github.com/jteeuwen/go-bindata/go-bindata
	go-bindata -nomemcopy -prefix='ui/dist' -o api/view/assets.go -pkg=view ./ui/dist/...

release: clean local
	mkdir -p bundles/$(PACKAGE_NAME)
	echo $(VERSION) > bundles/$(PACKAGE_NAME)/release.txt
	cd bundles && $(HASH) $(APP) > $(PACKAGE_NAME)/sha1.txt
	mv bundles/$(APP) bundles/$(PACKAGE_NAME)
	cd bundles && tar zcvf $(PACKAGE_NAME).tgz $(PACKAGE_NAME)

clean:
	rm -f api/view/assets.go
	rm -rf bundles/*

only-image:
	docker build -t ${IMAGE}:${GIT_COMMIT} .

dev:
	docker run --rm -it \
	 --name $(APP)-dev \
	 -p 8080:8080 \
	 -v $(PWD)/..:/opt/gopath/src/$(PKG)/.. \
	 -w /opt/gopath/src/$(PKG) \
	 $(GO_IMAGE) sh

dev-ui:
	docker run --rm -it \
	 --name $(APP)-dev \
	 -p 8080:8080 \
	 -v $(PWD)/ui:/opt/diego \
	 -w /opt/diego \
	 $(UI_IMAGE) bash
