FROM golang:1.22-alpine as builder

ARG APP_DIR=/src \
    LOCAL_PATH=./src

WORKDIR $APP_DIR

COPY ${LOCAL_PATH}/go.mod ${APP_DIR}/
COPY ${LOCAL_PATH}/go.sum ${APP_DIR}/
RUN go mod download

RUN apk add --no-cache git sqlite-dev gcc musl-dev
RUN go build -o main .

CMD ["/bin/bash", "./entrypoint.sh"]
