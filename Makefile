create_db:
	docker exec -it carbon_offsets_db psql -U docker -c "CREATE DATABASE carbon_offsets;"

migrate:
	docker exec -i carbon_offsets_db psql -d carbon_offsets -U docker -f /schema.sql

test:
	go test -v ./...

test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out