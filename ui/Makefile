PWD := $(shell pwd)
APP := diego-ui

UI_IMAGE := ckeyer/dev:vue

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
