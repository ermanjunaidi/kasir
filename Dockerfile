FROM golang:1.21-alpine

WORKDIR /app

# Copy seluruh isi project sekaligus (termasuk go.mod, go.sum, main.go, dll)
COPY . .

# Download dependensi
RUN go mod tidy

# Build binary
RUN go build -o main main.go

# Jalankan binary
CMD ["./main"]
