# Build image
FROM golang:1.22.1-alpine3.19 as builder

ENV GOPROXY=http://proxy.golang.org,direct

WORKDIR /app

COPY . .

RUN apk add \
    --no-cache \
    --allow-untrusted \
    --repository http://dl-cdn.alpinelinux.org/alpine/v3.19/main \
    --update git bash build-base

RUN go mod download

EXPOSE ${APP_PORT}

CMD [ "go", "run", "main.go" ]