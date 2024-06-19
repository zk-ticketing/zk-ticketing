FROM golang:1.21 AS build
WORKDIR /go/src
COPY openapi ./openapi
COPY service ./service
COPY repos ./repos
COPY main.go .
COPY go.sum .
COPY go.mod .

ENV CGO_ENABLED=0

RUN go build -o app .

FROM scratch AS runtime
COPY --from=build /go/src/app ./
EXPOSE 8080/tcp
ENTRYPOINT ["./app"]
