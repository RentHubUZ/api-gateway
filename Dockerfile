FROM golang:1.23.2 AS builder

WORKDIR /app

COPY . .

RUN go mod download

COPY .env .

RUN CGO_ENABLED=0 GOOS=linux go build -C ./cmd -a -installsuffix cgo -o ./../myapp .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/myapp .
COPY --from=builder /app/.env .
COPY --from=builder /app/internal/casbin/model.conf /app/internal/casbin/
COPY --from=builder /app/internal/casbin/policy.csv /app/internal/casbin/

EXPOSE 9996

CMD ["./myapp"]