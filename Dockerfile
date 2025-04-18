FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o server ./cmd/server

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server ./server
COPY config/app.yaml ./config/app.yaml
EXPOSE 8080
CMD ["./server"]
