FROM golang:1.19-alpine

WORKDIR /app

COPY ./ ./
RUN go mod tidy

RUN go build -o /service cmd/service.go

EXPOSE 10001

CMD [ "/service" ]