# build ui
# https://github.com/dokr/dev-base/blob/vue/Dockerfile
FROM ckeyer/dev:vue AS uidist

COPY ui /opt/diego

RUN cd /opt/diego ;\
	npm install ;\
	npm run build

# build server binary
# https://github.com/dokr/go/blob/1.10/Dockerfile
FROM ckeyer/dev:go AS gobin

ENV CGO_ENABLED=0

COPY . /go/src/github.com/ckeyer/diego
COPY --from=uidist /opt/diego/dist /go/src/github.com/ckeyer/diego/ui/dist

WORKDIR /go/src/github.com/ckeyer/diego

RUN make build

# ckeyer/diego
FROM alpine:edge

MAINTAINER Chuanjian Wang <me@ckeyer.com>

RUN apk add --update ca-certificates

COPY --from=gobin /go/src/github.com/ckeyer/diego/bundles/diego /usr/local/bin/

CMD ["/usr/local/bin/diego"]
