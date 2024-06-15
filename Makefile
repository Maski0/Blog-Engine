postgres:
	docker run --name Postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=Wasdqe@123 -d postgres:16-alpine

createdb:
	docker exec -it Postgres16 createdb --username=root --owner=root simple_blog
    
dropdb:
	docker exec -it Postgres16 dropdb simple_blog 

migrateup:
	migrate -path db/migration -database "postgresql://root:Wasdqe@123@localhost:5432/simple_blog?sslmode=disable" -verbose up 

migratedown:
	migrate -path db/migration -database "postgresql://root:Wasdqe@123@localhost:5432/simple_blog?sslmode=disable" -verbose down 

sqlc:
	sqlc generate 

.PHONY: postgres createdb dropdb migrateup migratedown sqlc