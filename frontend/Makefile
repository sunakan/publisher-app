export DOCKER_NODE_TAG=12-slim
export TINI_VERSION=v0.18.0
export APP_DOCKER_REPOSITORY=izukom/publisher-frontend
export APP_ENV=development

.PHONY: test
test:
	echo env = ${APP_ENV}
	echo env = ${TINI_VERSION}

.PHONY: build
build:
	docker build . \
		--build-arg DOCKER_NODE_TAG=${DOCKER_NODE_TAG} \
		--build-arg TINI_VERSION=${TINI_VERSION} \
		--target ${APP_ENV} \
		--tag ${APP_DOCKER_REPOSITORY}:${APP_ENV}

.PHONY: bash
bash:
	docker run \
		--rm \
		--interactive \
		--tty \
		--publish 3000:3000 \
		${APP_DOCKER_REPOSITORY}:${APP_ENV} \
		bash
