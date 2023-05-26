FROM golang:1.20 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o server ./cmd/server/server.go

FROM scratch
WORKDIR /server
COPY --from=builder /app/server .
CMD ["./server"]
