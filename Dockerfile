FROM golang:1.22 AS builder

WORKDIR /app

COPY src .

# Explicitly build for linux/amd64 architecture
RUN go build -o main .

FROM gcr.io/distroless/base

COPY --from=builder /app/main .

COPY --from=builder /app/static ./static

EXPOSE 8080

CMD ["./main"]

