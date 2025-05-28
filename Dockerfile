# Stage 1: Build the Go backend
FROM golang:1.21 AS builder

WORKDIR /app

# Copy go.mod and download dependencies
COPY backend/go.mod ./backend/
WORKDIR /app/backend
RUN go mod download

# Now go back and copy all source files including frontend
WORKDIR /app
COPY backend/ ./backend/
COPY frontend/ ./frontend/

# Build the Go app (output inside backend/)
WORKDIR /app/backend
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server main.go



# Stage 2: Lightweight runtime image
FROM alpine:latest

RUN apk add --no-cache ca-certificates

# Set working directory to /app/backend (where server will run from)
WORKDIR /app/backend

# Copy the server binary from builder
COPY --from=builder /app/backend/server .

# Copy frontend to sibling folder
COPY --from=builder /app/frontend /app/frontend

# Expose port
EXPOSE 8080

# Run server
CMD ["./server"]
