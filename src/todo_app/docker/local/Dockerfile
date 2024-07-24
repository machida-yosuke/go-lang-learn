FROM golang:1.22-alpine as builder

ARG APP_DIR=/src \
    LOCAL_PATH=./src

WORKDIR $APP_DIR

COPY ${LOCAL_PATH}/go.mod ${APP_DIR}/
COPY ${LOCAL_PATH}/go.sum ${APP_DIR}/
RUN go mod download

RUN apk add --no-cache git sqlite-dev gcc musl-dev

CMD ["go", "run", "main.go"]