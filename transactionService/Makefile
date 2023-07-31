.PHONY: proto
proto:
	mkdir -p pkg
	protoc --go_out=pkg  \
			dto/proto/transaction_v1/transaction.proto \
			dto/proto/user_v1/user.proto

migrations-up:
	migrate -path ./migrations/schema -database 'postgres://postgres:password123@localhost:5432/postgres?sslmode=disable' up

migrations-down:
	migrate -path ./migrations/schema -database 'postgres://postgres:password123@localhost:5432/postgres?sslmode=disable' down

run:
	docker-compose up

build_run:
	docker-compose up --build

init:
	go mod init

tidy:
	go mod tidy