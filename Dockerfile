FROM golang:1.20-alpine as builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o /gymshark

EXPOSE 8081

CMD ["/gymshark"]