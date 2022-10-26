FROM golang:alpine AS builder

LABEL maintaner="M Fitrah Ramadhan <mfitrahrmd420>"

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .
COPY .env .

RUN go get -d -v ./...

RUN go install -v ./...

RUN go build -o build .

FROM alpine:latest

WORKDIR /root

COPY --from=builder /app .
COPY --from=builder /app/.env .

EXPOSE 8001

CMD ["./build"]