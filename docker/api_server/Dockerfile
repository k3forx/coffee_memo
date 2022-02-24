FROM golang:1.17.7 AS builder
WORKDIR /app
COPY ./go/go.mod ./go/go.sum ./
RUN go mod download
COPY ./go .
RUN CGO_ENABLED=0 GOARCH=amd64 go build -o /go/bin/api_server ./cmd/api_server

FROM gcr.io/distroless/base-debian10
COPY --from=builder /go/bin/api_server /go/bin/api_server
EXPOSE 8000
ENTRYPOINT ["/go/bin/api_server"]
