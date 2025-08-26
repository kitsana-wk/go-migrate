DB_URL=postgres://myuser:mypassword@localhost:5432/mydb?sslmode=disable

migrate-up:
	migrate -path migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" down 1

migrate-force:
	migrate -path migrations -database "$(DB_URL)" force 1

new-migration:
	migrate create -ext sql -dir migrations -seq $(name)
