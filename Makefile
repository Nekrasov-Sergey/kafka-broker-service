PWD = ${CURDIR}
NAME = kafka-broker-service

# Запустить проект
.PHONY: run
run:
	go run $(PWD)/cmd/$(NAME)/

# Запустить docker
.PHONY: docker
docker:
	docker compose up -d