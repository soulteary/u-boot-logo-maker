FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build .

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/logo.png /
COPY --from=builder /app/u-boot-logo-maker /usr/bin/
ENTRYPOINT ["u-boot-logo-maker"]