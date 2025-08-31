.PHONY: build run test lint clean vet tidy swag mod-download

build: clean tidy swag
	@echo "Создаем Foodix..."
	go build -o ./bin/foodix ./cmd/server/main.go

run: build
	@echo "Запускаем Foodix..."
	./bin/foodix

test:
	@echo "Проводим тесты..."
	go test ./... -cover

lint:
	@echo "Запускаем golangci-lint..."
	golangci-lint run ./...

clean:
	@echo "Уборка..."
	go clean

vet:
	@echo "Запускаем go vet..."
	go vet

tidy:
	@echo "Запускаем go mod tidy..."
	go mod tidy

swag:
	@echo "Создаем документы swagger..."
	swag init --dir ./cmd/server/,./

mod-download:
	@echo "Загружаем модули Go..."
	go mod download

docker-build: build
	docker build -t foodix:latest -f Dockerfile .



# Имя пользователя GitHub (замени на свой)
USERNAME = skharchenko87
# Название образа
IMAGE = foodix
# Версия
VERSION = 1.0.0
# Регистр GHCR
REGISTRY = ghcr.io
# Полное имя образа
IMAGE_NAME = $(REGISTRY)/$(USERNAME)/$(IMAGE):$(VERSION)

.PHONY: ghcr-login ghcr-build ghcr-push

ghcr-login:
	@echo "===> Логинимся в GHCR..."
	@powershell -Command "echo $$env:GHCR_TOKEN | docker login $(REGISTRY) -u $(USERNAME) --password-stdin"

ghcr-build:
	@echo "===> Собираем образ..."
	docker build -t $(IMAGE_NAME) .

ghcr-push:
	@echo "===> Публикуем образ..."
	docker push $(IMAGE_NAME)

ghcr-publish: ghcr-login ghcr-build ghcr-push
	@echo "===> Образ опубликован $(IMAGE_NAME)"