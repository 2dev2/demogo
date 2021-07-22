

#if you want to run client and server seperate, then use client and server
#if client and server on same server address then use app

.PHONY: client
client:
	cd web && npm run start

.PHONY: server
server:
	go run ./bin/main.go

.PHONY: app
app:
	cd web && npm run build
	go run ./bin/main.go


