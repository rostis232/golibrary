up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "All Docker images are started"

up_build:
	@echo "Building and starting Docker images..."
	docker-compose up --build -d
	@echo "All Docker images are started"

down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

restart: down up