FROM golang:1.22 AS builder
RUN touch /empty

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

FROM golang:1.22

WORKDIR /app/

COPY --from=builder /app/myapp .
COPY --from=builder /empty .env

EXPOSE 8080

CMD ["./myapp"]
