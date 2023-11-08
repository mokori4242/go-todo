FROM golang:1.20

RUN mkdir -p /usr/local/go/src/app/app

WORKDIR /usr/local/go/src//app

RUN apt-get update && apt-get install -y git
RUN go install github.com/cosmtrek/air@latest

EXPOSE 8080
CMD ["air", "-c", ".air.toml"]