docker run --rm -v %cd%:/app -w /app golang:1.22 go mod init myapp
docker run --rm -v %cd%:/app -w /app golang:1.22 go mod tidy

## run make
###Local Go (ไม่ใช้ Docker)
```bash
go build   # สร้าง binary
go run     # รันโปรแกรม
go test    # รันทดสอบ
go lint    # เช็กโค้ด
```

docker-compose logs -f app

## docker run
```bash
docker-compose up
```

## down:
```bash
docker-compose down
```

## install CLI
```bash
brew install golang-migrate
migrate -version
```

go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate create -ext sql -dir migrations -seq create_users_table

docker-compose exec app sh
