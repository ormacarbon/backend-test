install:
	cd docker && docker compose up -d
	cd client && npm install

runapi:
	cd server && go run cmd/main.go

runclient:
	cd client && npm run dev