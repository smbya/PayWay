FROM golang:1.25.7-alpine AS builder

RUN apk update && apk upgrade && \
    apk --update add git make bash build-base

WORKDIR /app

COPY ./service_web ./service_web
COPY ./service_callback ./service_callback

RUN go build -o builded/web service_web/cmd/main.go
RUN go build -o builded/callback service_callback/cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/builded /app

EXPOSE 80 81

CMD /app/web & /app/callback