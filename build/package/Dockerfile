FROM golang:1.14-alpine AS builder
WORKDIR /app/
ENV APP_ROOT /app
COPY . .
RUN apk add --update make && make build
#RUN go build -o bin/app cmd/main.go

FROM alpine:3.11
COPY --from=builder /app /app
RUN chmod +x /app/bin/*
CMD ["./app/bin/app"]