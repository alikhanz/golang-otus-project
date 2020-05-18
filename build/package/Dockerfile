FROM golang:1.14-stretch AS builder
WORKDIR /app/
ENV APP_ROOT /app
COPY . .
RUN make build

FROM alpine:3.11
COPY --from=builder /app /app
COPY build/package/wait-for /usr/bin/wait-for
RUN chmod -R +x /usr/bin/wait-for && \
    chmod -R +x /app/bin
CMD ["/app/bin/app"]