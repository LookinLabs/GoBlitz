# Build image
FROM golang:1.22.1-alpine3.19 as builder

WORKDIR /app

COPY . .

RUN apk add \
    --no-cache \
    --allow-untrusted \
    --repository https://dl-cdn.alpinelinux.org/alpine/v3.19/main \
    --update git bash build-base

RUN go mod download

EXPOSE ${APP_PORT}

CMD [ "go", "run", "main.go" ]