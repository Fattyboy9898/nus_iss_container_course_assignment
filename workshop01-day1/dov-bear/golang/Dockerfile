
FROM golang:1.22 as builder

WORKDIR /app

COPY . .
RUN go mod download

ENV CGO_ENABLED=0
RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY public ./public
COPY templates ./templates


EXPOSE 3000
ENTRYPOINT [ "./main" ]
CMD [ "port", "name", "hash" ]