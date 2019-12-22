PWD := $(shell pwd)
APP := diego-ui

UI_IMAGE := ckeyer/dev:vue

env:
	npm config set strict-ssl false
	npm config set registry https://registry.npm.taobao.org
	npm config set sass-binary-site http://npm.taobao.org/mirrors/node-sass
	npm install

build-in-docker:
	docker run --rm \
	 --name ${APP} \
	 -v ${PWD}:/tmp/${APP} \
	 -w /tmp/${APP} \
	 ${UI_IMAGE} make build-ui

build-ui: clean-ui
	npm install
	npm run build

clean-ui:
	@rm -rf ui/dist
	@rm -rf dist

dev-ui:
	docker run --rm -it \
	 --name $(APP)-dev \
	 -p 8080:8080 \
	 -v $(PWD):/opt/diego \
	 -w /opt/diego \
	 $(UI_IMAGE) bash
