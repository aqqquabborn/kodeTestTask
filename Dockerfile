FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY .env .env

EXPOSE 8080

RUN go build -o kodeTestTask

CMD ["./kodeTestTask"]
