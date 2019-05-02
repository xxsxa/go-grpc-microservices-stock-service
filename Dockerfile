FROM golang:alpine as builder
ENV GO111MODULE=on
RUN apk --no-cache add git
ADD . /app/stock-service
WORKDIR /app/stock-service
COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o stock-service


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/stock-service .

CMD ["./stock-service"]