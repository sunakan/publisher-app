ARG DOCKER_GO_TAG
ARG DOCKER_ALPINE_TAG

################################################################################
# Builder
################################################################################
FROM golang:${DOCKER_GO_TAG} as builder
RUN apk update \
  && apk add \
    git \
    make
WORKDIR /go/src/github.com/sunakan/publisher-app/backend/go-app/
COPY go-app/* ./
RUN go get \
  && go build

################################################################################
# Production
################################################################################
FROM alpine:${DOCKER_ALPINE_TAG} as production
ENV PORT=8080

RUN addgroup -S publisher \
  && adduser -S publisher -G publisher \
  && mkdir --parents /publisher/log \
  && chown --recursive publisher /publisher \
  && chgrp --recursive publisher /publisher

COPY --from=builder /go/src/github.com/sunakan/publisher-app/backend/go-app/go-app /usr/local/bin/publisher-api

USER publisher
ENTRYPOINT ["publisher-api"]
