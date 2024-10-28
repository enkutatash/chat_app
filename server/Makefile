postgresinit:
	docker run --name postgres -p 5433:5432 -e POSTGRES_PASSWORD=enku0811 -d postgres
postgres:
	docker exec -it postgres psql -U postgres
createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres goc
dropdb:
	docker exec -it postgres dropdb goc
migrateup:
	migrate -path db/migrate -database "postgresql://postgres:enku0811@localhost:5433/gochat?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrate -database "postgresql://postgres:enku0811@localhost:5433/gochat?sslmode=disable" -verbose down

.PHONEY: postgresinit postgres createdb dropdb migrateup migratedown