FROM golang:1.22-alpine as builder

WORKDIR /app

# Goモジュールをキャッシュするために必要なファイルをコピー
COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN apk add --no-cache git sqlite-dev gcc musl-dev
COPY . .
RUN go build -o main .

CMD ["./main"]