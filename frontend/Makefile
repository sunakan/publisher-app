export DOCKER_NODE_TAG=12-slim
export TINI_VERSION=v0.18.0
export APP_DOCKER_REPOSITORY=izukom/publisher-frontend
export APP_ENV=development

.PHONY: build
build:
	docker build . \
		--build-arg DOCKER_NODE_TAG=${DOCKER_NODE_TAG} \
		--build-arg TINI_VERSION=${TINI_VERSION} \
		--target ${APP_ENV} \
		--tag ${APP_DOCKER_REPOSITORY}:${APP_ENV}

.PHONY: build-production
build-production:
	make build APP_ENV=production

.PHONY: bash
bash:
	docker run \
		--rm \
		--interactive \
		--tty \
		--publish 3000:3000 \
		--user `id --user`:`id --group` \
		--mount type=bind,source=${PWD}/app/,target=/home/node/app/ \
		${APP_DOCKER_REPOSITORY}:${APP_ENV} \
		bash

################################################################################
# デプロイ
# TODO ファイル分割
################################################################################
export HEROKU_APP_NAME=izukom-publisher-frontend
.PHONY: h-login
h-login:
	heroku login --interactive
.PHONY: h-container-login
h-container-login:
	heroku container:login
.PHONY: h-create-app
h-create-app:
	heroku apps | grep ${HEROKU_APP_NAME} || heroku create ${HEROKU_APP_NAME}
.PHONY: h-add-app-for-git
h-add-app-for-git:
	heroku git:remote --app ${HEROKU_APP_NAME}
.PHONY: h-build-production
h-build-production: h-create-app
	docker tag ${APP_DOCKER_REPOSITORY}:production registry.heroku.com/${HEROKU_APP_NAME}/web
.PHONY: h-push
h-push: h-build-production
	docker push registry.heroku.com/${HEROKU_APP_NAME}/web
	docker image rm registry.heroku.com/${HEROKU_APP_NAME}/web
.PHONY: h-release
h-release:
	heroku container:release web
