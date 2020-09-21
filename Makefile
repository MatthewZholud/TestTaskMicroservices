run:
	docker-compose up
	cd cmd/api && go run main.go
migration:
	cd migrator && docker build -t migrator . && docker run --network host migrator -path=/migrations/ -database "postgresql://postgres:mypassword@localhost:5432/company_manager?sslmode=disable" up
