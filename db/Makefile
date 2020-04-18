export DOCKER_MIGRATE_TAG=v4.10.0
export DOCKER_POSTGRES_TAG=12
export POSTGRESQL_URL="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:5432/${DB_NAME}"

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

select-all:
	docker run \
		--rm \
		--interactive \
		--tty \
		--hostname pg \
		--name pg \
		--mount type=bind,source=${PWD}/,target=/home/pg/ \
		--workdir /home/pg/ \
		--env PGPASSWORD=${DB_PASSWORD} \
		postgres:${DOCKER_POSTGRES_TAG} \
		bash -c \
			'psql --command "SELECT * FROM publisher;" --host ${DB_HOST} --port 5432 --username ${DB_USER} --dbname ${DB_NAME}'

insert-test:
	docker run \
		--rm \
		--interactive \
		--tty \
		--hostname pg \
		--name pg \
		--mount type=bind,source=${PWD}/,target=/home/pg/ \
		--workdir /home/pg/ \
		--env PGPASSWORD=${DB_PASSWORD} \
		postgres:${DOCKER_POSTGRES_TAG} \
		bash -c \
			"psql --command \"INSERT INTO publisher (name) VALUES ('技術評論社');\" --host ${DB_HOST} --port 5432 --username ${DB_USER} --dbname ${DB_NAME}"

truncate:
	docker run \
		--rm \
		--interactive \
		--tty \
		--hostname pg \
		--name pg \
		--mount type=bind,source=${PWD}/,target=/home/pg/ \
		--workdir /home/pg/ \
		--env PGPASSWORD=${DB_PASSWORD} \
		postgres:${DOCKER_POSTGRES_TAG} \
		bash -c \
			"psql --command \"TRUNCATE TABLE publisher;\" --host ${DB_HOST} --port 5432 --username ${DB_USER} --dbname ${DB_NAME}"

bash:
	docker run \
		--rm \
		--interactive \
		--tty \
		--hostname pg \
		--name pg \
		--mount type=bind,source=${PWD}/,target=/home/pg/ \
		--workdir /home/pg/ \
		--env PGPASSWORD=${DB_PASSWORD} \
		postgres:${DOCKER_POSTGRES_TAG} \
		bash

puml:
	docker run \
		--rm \
		--detach \
		--publish 8080:8080 \
		plantuml/plantuml-server:jetty
