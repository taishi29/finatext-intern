FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# デフォルトは何もしない（Makefileで実行）
CMD ["go", "run", "cmd/server/main.go"]
