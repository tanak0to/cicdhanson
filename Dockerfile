FROM golang:1.22.2 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o cicdhanson

FROM scratch
COPY --from=builder /app/cicdhanson /app
ENTRYPOINT ["/app"]
