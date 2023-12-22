FROM golang:latest

WORKDIR ~/Containers/back-end-container

COPY . .

EXPOSE 41750

RUN go build -o main .

CMd ["./main"]
