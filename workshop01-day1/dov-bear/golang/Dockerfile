
FROM golang:1.22

WORKDIR /app

COPY public ./public
COPY templates ./templates
COPY . .
COPY go.mod go.sum ./

RUN go mod download

RUN go build -o main .

EXPOSE 3000
CMD ["./main"]
