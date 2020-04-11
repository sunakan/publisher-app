export DOCKER_MIGRATE_TAG=v4.10.0
export DOCKER_POSTGRES_TAG=12
export POSTGRESQL_URL='postgres://yjnwdibd:dcDGXTZeORGq77WsBPkRZ9BvFAYiK8_N@arjuna.db.elephantsql.com:5432/yjnwdibd'

up:
	docker run \
		--rm \
		--interactive \
		--tty \
		--hostname migrate \
		--name migrate \
		--user `id --user`:`id --group` \
		--mount type=bind,source=${PWD}/,target=/home/migrate/ \
		--workdir /home/migrate/ \
		migrate/migrate:${DOCKER_MIGRATE_TAG} \
		-database ${POSTGRESQL_URL} -path ./migrations up

down:
	docker run \
		--rm \
		--interactive \
		--tty \
		--hostname migrate \
		--name migrate \
		--user `id --user`:`id --group` \
		--mount type=bind,source=${PWD}/,target=/home/migrate/ \
		--workdir /home/migrate/ \
		migrate/migrate:${DOCKER_MIGRATE_TAG} \
		-database ${POSTGRESQL_URL} -path ./migrations down

bash:
	docker run \
		--rm \
		--interactive \
		--tty \
		--hostname pg \
		--name pg \
		--mount type=bind,source=${PWD}/,target=/home/pg/ \
		--workdir /home/pg/ \
		--env PGPASSWORD=dcDGXTZeORGq77WsBPkRZ9BvFAYiK8_N \
		postgres:${DOCKER_POSTGRES_TAG} \
		bash

puml:
	docker run \
		--rm \
		--detach \
		--publish 8080:8080 \
		plantuml/plantuml-server:jetty

#
# psql --host arjuna.db.elephantsql.com --port 5432 --username yjnwdibd --dbname yjnwdibd
#
#----------
# SELECT *
# FROM publisher
#----------
