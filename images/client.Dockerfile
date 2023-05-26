FROM golang:1.20 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o client ./cmd/client/client.go

FROM scratch
WORKDIR /client
COPY --from=builder /app/client .
CMD ["./client"]
