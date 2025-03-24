# Dockerfile（プロジェクトルートに配置）
FROM golang:1.23.5

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

CMD ["./main"]
