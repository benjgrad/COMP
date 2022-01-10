FROM golang:1.16

WORKDIR /app
COPY go.mod ./

COPY . .

RUN go mod download

RUN go build -o /comp ./src/app

CMD ["/comp"]