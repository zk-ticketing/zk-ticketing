FROM golang:1.21-alpine as builder

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN go build -o main .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 3000
CMD ["./main"]
