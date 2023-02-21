migrateup:
	migrate -path db/migration -database "postgresql://xinqi:xinqisecret@localhost:5432/simplebankpsql?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://xinqi:xinqisecret@localhost:5432/simplebankpsql?sslmode=disable" -verbose down
sqlc:
	sqlc generate
.PHONY: migrateup migratedown slqc