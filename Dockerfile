FROM arm64v8/golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o main .

FROM arm64v8/debian:bullseye-slim

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

CMD [ "./main" ]