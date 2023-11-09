FROM golang:1.19

EXPOSE 8000

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /mini-server

CMD ["/mini-server"]
