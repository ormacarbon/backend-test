.PHONY: build run air stop logs db-up db-down clean

# Nome do container do banco de dados
DB_CONTAINER = postgres_competition

# 🔹 Subir apenas o banco de dados com docker-compose
db-up:
	docker-compose up -d db

# 🔹 Parar e remover o banco de dados
db-down:
	docker-compose down

# 🔹 Rodar a aplicação Go normalmente
run:
	go run main.go

# 🔹 Rodar a aplicação com live reload (precisa do Air instalado)
air:
	air

# 🔹 Parar o banco de dados e limpar containers
stop:
	docker stop $(DB_CONTAINER) || true
	docker rm $(DB_CONTAINER) || true

# 🔹 Ver logs do banco de dados
logs:
	docker logs -f $(DB_CONTAINER)

# 🔹 Limpar volumes e imagens não usadas
clean: stop db-down
	docker volume rm $(shell docker volume ls -q) || true
