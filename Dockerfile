FROM golang:1.18 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main cmd/main.go
FROM gcr.io/distroless/base-debian10
COPY --from=builder /app/main .
EXPOSE 8000
CMD ["/main"]
