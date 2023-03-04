FROM golang:latest

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app

COPY .. .
RUN go mod download

CMD ["air", "-c",  ".air.toml"]
