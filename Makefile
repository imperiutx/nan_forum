startdb:
	docker run --name nanf_db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.2-alpine

createdb:
	docker exec -it nanf_db createdb --username=root --owner=root nanf_db

dropdb:
	docker exec -it nanf_db dropdb nanf_db

mic:
	migrate create -ext sql -dir db/migration -seq init_schema
	
miup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/nanf_db?sslmode=disable" -verbose up

midown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/nanf_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

proto:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

.PHONY: proto