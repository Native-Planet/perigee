FROM docker.io/library/golang:1.23.2 AS builder
WORKDIR /app
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY . /app/
RUN mkdir -p /app/out/out
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o out/perigee main.go

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /app/out/ .
ENV PATH=$$PATH:/
EXPOSE 8080
ENTRYPOINT ["perigee", "serve"]