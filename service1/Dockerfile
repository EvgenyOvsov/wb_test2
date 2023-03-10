FROM golang:alpine as builder
RUN apk add make
WORKDIR /app
COPY . .
RUN make compile

FROM busybox:latest
WORKDIR /app
COPY --from=builder /app/bin/service ./service
ENTRYPOINT ["/app/service"]