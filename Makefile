install:
	docker compose up -d
	cd client && npm install

runapi:
	docker compose up -d
	cd server && go run cmd/main.go

runclient:
	cd client && npm run dev