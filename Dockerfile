FROM golang:1.22.1 as builder

WORKDIR /app

COPY go.mod go.sum .env apispec.yaml ./
RUN go mod download

COPY adapters .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./cmd/webserver/main ./cmd/webserver

FROM scratch

# Salin binary aplikasi dan file .env yang telah dibangun ke dalam image scratch
COPY --from=builder /app/cmd/webserver/main .
COPY --from=builder /app/.env .
COPY --from=builder /app/apispec.yaml .

EXPOSE 8081

# Jalankan aplikasi
CMD ["./cmd/webserver/main"]