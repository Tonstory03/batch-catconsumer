FROM golang:1.17-alpine as builder

WORKDIR /app/

COPY . .

RUN go get -d -v ./...

RUN go install -v ./... -o /batch-catconsumer

EXPOSE 8080

CMD ["batch-catconsumer"]