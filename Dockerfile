FROM golang:1.20-alpine as builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o /gymshark

FROM builder AS run-test-stage
RUN go test -v ./...

EXPOSE 8081

CMD ["/gymshark"]