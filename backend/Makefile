export TMP_HOME=${PWD}/tmp/tmp-home
export TMP_GOPATH_CACHE=${PWD}/tmp/GOPATH_CACHE

export DOCKER_GO_TAG=1.14-alpine
export DOCKER_ALPINE_TAG=3
export APP_DOCKER_REPOSITORY=izukom/publisher-backend

#.PHONY: setup-dev-env
#setup-dev-env:
#	mkdir --parents ${TMP_HOME}
#	mkdir --parents ${TMP_GOPATH_CACHE}

.PHONY: build
build:
	docker build ./ \
		--build-arg DOCKER_GO_TAG=${DOCKER_GO_TAG} \
		--target builder \
		--tag ${APP_DOCKER_REPOSITORY}:latest

.PHONY: build-production
build-production:
	docker build ./ \
		--build-arg DOCKER_GO_TAG=${DOCKER_GO_TAG} \
		--build-arg DOCKER_ALPINE_TAG=${DOCKER_ALPINE_TAG} \
		--target production \
		--tag ${APP_DOCKER_REPOSITORY}:production

.PHONY: ash
ash:# setup-dev-env
	docker run \
		--rm \
		--interactive \
		--tty \
		--hostname publisher-backend \
		--name publisher-backend \
		--publish 8080:8080 \
		--env PORT=8080 \
		--env HOME=/usr/local/tmp-home/ \
		--mount type=bind,source=${PWD}/go-app/,target=/go/src/github.com/sunakan/publisher-app/backend/go-app/ \
		--workdir /go/src/github.com/sunakan/publisher-app/backend/go-app/ \
		${APP_DOCKER_REPOSITORY}:latest \
		ash

#export APP_URL=http://127.0.0.1:8080
export APP_URL=https://publisher-api.herokuapp.com
.PHONY: post
post:
	curl \
		--request POST \
		--header "Content-Type: application/json" \
		--data '{"name":"達人出版"}' \
		${APP_URL}/publisher/
	curl \
		--request POST \
		--header "Content-Type: application/json" \
		--data '{"name":"技術評論社"}' \
		${APP_URL}/publisher/
	curl \
		--request POST \
		--header "Content-Type: application/json" \
		--data '{"name":"オーム社"}' \
		${APP_URL}/publisher/
	curl \
		--request POST \
		--header "Content-Type: application/json" \
		--data '{"name":"インプレス"}' \
		${APP_URL}/publisher/
	curl \
		--request POST \
		--header "Content-Type: application/json" \
		--data '{"name":"ラムダノート"}' \
		${APP_URL}/publisher/
	curl \
		--request POST \
		--header "Content-Type: application/json" \
		--data '{"name":"翔泳社"}' \
		${APP_URL}/publisher/
	curl \
		--request POST \
		--header "Content-Type: application/json" \
		--data '{"name":"秀和システム"}' \
		${APP_URL}/publisher/
	curl \
		--request POST \
		--header "Content-Type: application/json" \
		--data '{"name":"オライリージャパン"}' \
		${APP_URL}/publisher/
	curl \
		--request POST \
		--header "Content-Type: application/json" \
		--data '{"name":"KADOKAWA"}' \
		${APP_URL}/publisher/

# デバッグ用
#		--mount type=bind,source=${TMP_HOME}/,target=/usr/local/tmp-home/ \
#		--mount type=bind,source=${TMP_GOPATH_CACHE}/,target=/go/src/ \
#		--user `id -u`:`id -g` \

################################################################################
# デプロイ
# TODO ファイル分割
################################################################################
export HEROKU_APP_NAME=publisher-api
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
h-build-production: h-create-app build-production
	docker tag ${APP_DOCKER_REPOSITORY}:production registry.heroku.com/${HEROKU_APP_NAME}/web
.PHONY: h-push
h-push: h-build-production
	docker push registry.heroku.com/${HEROKU_APP_NAME}/web
.PHONY: h-release
h-release:
	heroku container:release web
