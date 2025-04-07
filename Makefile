install:
	docker compose up -d
	cd client && npm install

api:
	cd server && go run cmd/main.go

cli:
	cd client && npm run dev