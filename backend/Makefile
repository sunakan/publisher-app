export TMP_HOME=${PWD}/tmp/tmp-home
export TMP_GOPATH_CACHE=${PWD}/tmp/GOPATH_CACHE

export DOCKER_GO_TAG=1.14-alpine
export APP_DOCKER_REPOSITORY=izukom/publisher-backend

.PHONY: setup-dev-env
setup-dev-env:
	mkdir --parents ${TMP_HOME}
	mkdir --parents ${TMP_GOPATH_CACHE}

.PHONY: build
build:
	docker build ./ \
		--build-arg DOCKER_GO_TAG=${DOCKER_GO_TAG} \
		--tag ${APP_DOCKER_REPOSITORY}:latest

.PHONY: ash
ash: setup-dev-env
	docker run \
		--rm \
		--interactive \
		--tty \
		--user `id -u`:`id -g` \
		--hostname publisher-backend \
		--name publisher-backend \
		--publish 8080:8080 \
		--env PORT=8080 \
		--env HOME=/usr/local/tmp-home/ \
		--mount type=bind,source=${TMP_HOME}/,target=/usr/local/tmp-home/ \
		--mount type=bind,source=${TMP_GOPATH_CACHE}/,target=/go/src/ \
		--mount type=bind,source=${PWD}/go-app/,target=/usr/local/go-app/ \
		--workdir /usr/local/go-app/ \
		${APP_DOCKER_REPOSITORY}:latest \
		ash

.PHONY: clean-complete
clean-complete:
	rm -rf tmp/
