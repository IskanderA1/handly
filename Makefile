postgres:
	docker run --name handly-pg -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root250700 -d postgres:12

createdb:
	docker exec -it handly-pg createdb --username=root --owner=root handly

dropdb:
	docker exec -it handly-pg dropdb handly

migrateup:
	migrate -path iternal/db/migrations -database "postgresql://root:root250700@localhost:5432/handly?sslmode=disable" -verbose up

migratedown:
	migrate -path iternal/db/migrations -database "postgresql://root:root250700@localhost:5432/handly?sslmode=disable" -verbose down

migrateupone:
	migrate -path iternal/db/migrations -database "postgresql://root:root250700@localhost:5432/handly?sslmode=disable" -verbose up 1

migratedownone:
	migrate -path iternal/db/migrations -database "postgresql://root:root250700@localhost:5432/handly?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...
	
server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown migrateupone migratedownone sqlc test server mock