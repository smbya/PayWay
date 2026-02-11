FROM golang:1.25.7-alpine AS builder

RUN apk update && apk upgrade && \
    apk --update add git make bash build-base

WORKDIR /app

COPY ./app .

RUN go build -o /app_build/worker cmd/worker/main.go
RUN go build -o /app_build/projection_worker cmd/projection_worker/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app_build /app

CMD /app/worker & /app/projection_worker