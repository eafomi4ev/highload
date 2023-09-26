
docker_build:
	docker build -t otus_highload:v1.0.0 .

# POSTGRES_HOST=host.docker.internal - для доступа из контейнера к БД на localhost на MacOS
docker_run:
	docker run \
	-p 8000:8000 \
	-e POSTGRES_HOST=host.docker.internal \
	-e POSTGRES_PORT=5432 \
	-e POSTGRES_USER_NAME=postgres \
	-e POSTGRES_DB_PASSWORD= \
	-e POSTGRES_DB_NAME=social_db \
	otus_highload:v1.0.0
