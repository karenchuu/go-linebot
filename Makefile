up: ## Start all containers in background
	docker-compose up -d

down: ## stop and clean all data
	docker-compose down

restart: ## Restart all containers
	docker-compose down
	docker-compose up -d

stop: ## Stop all containers
	docker-compose stop

logs: ## Show logs for all containers
	docker-compose logs --tail=100
