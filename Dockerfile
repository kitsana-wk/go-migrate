FROM golang:1.25.0-alpine3.22

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.* ./
# COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]