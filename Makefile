PWD := $(shell pwd)
APP := diego
PKG := github.com/ckeyer/$(APP)

GO := CGO_ENABLED=0 GOBIN=${PWD}/bundles go

# OS := $(shell go env GOOS)
# ARCH := $(shell go env GOARCH)
# VERSION := $(shell cat VERSION.txt)
# GIT_COMMIT := $(shell git rev-parse --short HEAD)
# GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
# BUILD_AT := $(shell date "+%Y-%m-%dT%H:%M:%SZ%z")
# PACKAGE_NAME := $(APP)$(VERSION).$(OS)-$(ARCH)

LD_FLAGS := -X github.com/ckeyer/commons/version.version=$(VERSION) \
 -X github.com/ckeyer/commons/version.gitCommit=$(GIT_COMMIT) \
 -X github.com/ckeyer/commons/version.buildAt=$(BUILD_AT) -w

DEV_IMAGE := golang:alpine
UIDEV_IMAGE := registry.cn-beijing.aliyuncs.com/ckeyer/dev:vue

gorun:
	$(GO) run -ldflags="$(LD_FLAGS)" main.go

local:
	$(GO) install -a -ldflags="$(LD_FLAGS)" .
	make hash

build:
	tools/build.sh
	# $(GO) build -a -ldflags="$(LD_FLAGS)" -o bundles/$(APP) main.go
	# make hash

hash:
	$(HASH) bundles/$(APP)

test:
	$(GO) test $$(go list ./... |grep -v "vendor")

test-in-docker:
	docker run --rm \
	 -v ${PWD}:/go/src/${PKG} \
	 -w /go/src/${PKG} \
	 golang:alpine \
	 go test -ldflags="$(LD_FLAGS)" $$(go list ./... |grep -v "vendor")

go-bindata:
	which go-bindata || go get github.com/jteeuwen/go-bindata/go-bindata
	go-bindata -nomemcopy -prefix='ui/dist' -o api/view/assets.go -pkg=view ./ui/dist/...

release: clean local
	mkdir -p bundles/$(PACKAGE_NAME)
	echo $(VERSION) > bundles/$(PACKAGE_NAME)/release.txt
	cd bundles && $(HASH) $(APP) > $(PACKAGE_NAME)/sha1.txt
	mv bundles/$(APP) bundles/$(PACKAGE_NAME)
	cd bundles && tar zcvf $(PACKAGE_NAME).tgz $(PACKAGE_NAME)

clean:
	rm -rf bundles/*

only-image:
	docker build -t ${IMAGE}:${GIT_COMMIT} .

dev:
	docker run --rm -it \
	 --name $(APP)-dev \
	 -p 8080:8080 \
	 -v $(PWD)/..:/opt/gopath/src/$(PKG)/.. \
	 -w /opt/gopath/src/$(PKG) \
	 $(DEV_IMAGE) sh

dev-ui:
	docker run --rm -it \
	 --name $(APP)-dev \
	 -p 8080:8080 \
	 -v $(PWD)/ui:/opt/ui \
	 -w /opt/ui \
	 $(UIDEV_IMAGE) bash
