gen:
	protoc --proto_path=cmd/api/proto cmd/api/proto/*.proto --go_out=plugins=grpc:cmd/api/
clean:
	rm cmd/api/pb/*.go 
run:
	go run tms/cmd/api main.go

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

mysql:
	docker run --name mysql8 -p 3306:3306  -e MYSQL_ROOT_PASSWORD=secret -d mysql:8

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root chat

dropdb:
	docker exec -it postgres12 dropdb chat

# migrateup:
# 	migrate -path cmd/model/sql/migration -database "postgresql://root:secret@localhost:5432/chat?sslmode=disable" -verbose up

# migrateup1:
# 	migrate -path cmd/model/sql/migration -database "postgresql://root:secret@localhost:5432/chat?sslmode=disable" -verbose up 1

# migratedown:
# 	migrate -path cmd/model/sql/migration -database "postgresql://root:secret@localhost:5432/chat?sslmode=disable" -verbose down

# migratedown1:
# 	migrate -path cmd/model/sql/migration -database "postgresql://root:secret@localhost:5432/chat?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run cmd/api/main.go


.PHONY: gen postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server
