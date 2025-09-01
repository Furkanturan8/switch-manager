FROM golang:1.21-alpine AS builder

WORKDIR /app

# Go modüllerini kopyala
COPY go.mod ./
RUN go mod download

# Kaynak kodları kopyala
COPY . .

# Uygulamayı derle
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

# Final image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Derlenen binary'yi kopyala
COPY --from=builder /app/main .

# Port'u expose et
EXPOSE 8080

# Uygulamayı çalıştır
CMD ["./main"]
