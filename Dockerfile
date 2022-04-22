# syntax=docker/dockerfile:1

FROM golang:alpine as builder

WORKDIR /build

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

# Build with compile-time parameters
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o server .

# Get clean app bin
FROM scratch
COPY --from=builder /build/server /app/
WORKDIR /app

EXPOSE 1323

CMD ["./server"]