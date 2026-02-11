FROM golang:1.25.7-alpine AS builder

RUN apk update && apk upgrade && \
    apk --update add git make bash build-base

WORKDIR /app

COPY ./app .

RUN go build -o /app_build/web cmd/web/main.go
RUN go build -o /app_build/callback cmd/callback/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app_build /app

EXPOSE 80 81

CMD /app/web & /app/callback