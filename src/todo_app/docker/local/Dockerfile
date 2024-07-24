FROM golang:1.22-alpine as builder

ARG APP_DIR=/src \
    LOCAL_PATH=./src

WORKDIR $APP_DIR

RUN go install github.com/air-verse/air@latest

COPY ${LOCAL_PATH}/go.mod ${APP_DIR}/
COPY ${LOCAL_PATH}/go.sum ${APP_DIR}/
RUN go mod download

RUN apk upgrade --update && apk --no-cache add git gcc musl-dev


CMD ["air", "-c", ".air.toml"]