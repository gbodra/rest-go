default:
	@echo "=============Building Local API============="
	docker build -f ./Dockerfile -t docker-rest-go .

up: default
	@echo "=============Starting API Locally============="
	docker-compose up -d

logs:
	docker-compose logs -f

down:
	docker-compose down

test:
	go test -v -cover ./...

clean: down
	@echo "=============Cleaning Up============="
	rm -f api
	docker system prune -f
	docker volume prune -f