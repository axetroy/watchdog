# builder for frontend
FROM node:14-alpine AS node-builder

WORKDIR /app

COPY ./web ./

RUN yarn

RUN yarn build

# builder for backend
FROM golang:1.16.2-alpine AS go-builder

WORKDIR /app

COPY --from=node-builder /app/dist ./web/dist

COPY job.go alarm.go serve.go stroage.go config.go go.mod go.sum ./
COPY ./vendor ./vendor
COPY ./internal ./internal
COPY ./cmd/watchdog/main.go ./cmd/watchdog/main.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -trimpath -asmflags=-trimpath=$GOPATH -ldflags "-s -w" -o ./bin/watchdog ./cmd/watchdog/main.go

# target
FROM alpine:3.13
WORKDIR /app
COPY --from=go-builder /app/bin .

ENV PORT=80

EXPOSE 80

CMD ["./watchdog", "./watchdog.config.json"]