FROM golang:1.22 AS builder

WORKDIR /app

# Caching dependencies
COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2
FROM gcr.io/distroless/base

COPY --from=builder /app/main .
COPY --from=builder /app/static ./static

EXPOSE 8080

ENTRYPOINT ["./main"]
