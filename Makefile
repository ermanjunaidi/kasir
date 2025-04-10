APP_NAME = main

.PHONY: run build clean tidy test env

run:
	@echo "🔧 Running $(APP_NAME)..."
	go run $(APP_NAME).go

build:
	@echo "📦 Building $(APP_NAME)..."
	go build -o $(APP_NAME) $(APP_NAME).go

clean:
	@echo "🧹 Cleaning up..."
	rm -f $(APP_NAME)

tidy:
	@echo "🧼 Tidying Go modules..."
	go mod tidy

test:
	@echo "🧪 Running tests..."
	go test ./...

env:
	@echo "🌱 Exporting environment variables..."
	export $$(cat .env | xargs)

IMAGE_NAME = user-api
TAG = latest

docker:
	@echo "🐳 Building Docker image $(IMAGE_NAME):$(TAG)..."
	docker build -t $(IMAGE_NAME):$(TAG) .
