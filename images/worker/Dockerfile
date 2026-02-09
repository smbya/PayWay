FROM golang:1.25.7-alpine AS builder

RUN apk update && apk upgrade && \
    apk --update add git make bash build-base

WORKDIR /app

COPY ./service_worker ./service_worker
COPY ./service_projection_worker ./service_projection_worker

RUN go build -o builded/worker service_worker/cmd/main.go
RUN go build -o builded/projection_worker service_projection_worker/cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/builded /app

CMD /app/worker & /app/projection_worker