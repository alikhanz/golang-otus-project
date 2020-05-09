FROM golang:1.14-stretch AS builder
WORKDIR /app/
ENV APP_ROOT /app
COPY . .

FROM alpine:3.11
COPY --from=builder /app /app
RUN chmod +x /app/bin/*
CMD ["./app/bin/app"]