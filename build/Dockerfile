FROM golang:1.22 AS builder

WORKDIR /app

COPY . ./
RUN go mod download


RUN go build -o forum ./cmd/main.go
FROM debian:stable-slim
WORKDIR /app
COPY --from=builder /app/forum .
COPY --from=builder /app/internal/data ../internal/data
COPY --from=builder app/internal/presentation/templates ../internal/presentation/templates
COPY  --from=builder app/internal/presentation/static ../internal/presentation/static
EXPOSE 8080

CMD ["./forum"]
