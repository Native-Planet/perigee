FROM docker.io/library/golang:1.23.2 AS builder
WORKDIR /app
COPY perigee/go.mod go.mod
COPY perigee/go.sum go.sum
RUN go mod download
COPY perigee/ /app/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app main.go
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /app/app .
EXPOSE 8080
ENTRYPOINT ["/app"]