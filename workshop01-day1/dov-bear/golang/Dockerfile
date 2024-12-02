
FROM golang:1.22

WORKDIR /app

COPY public ./public
COPY templates ./templates
COPY . .
RUN go mod download

RUN go build -o main .

EXPOSE 3000
ENTRYPOINT [ "./main" ]
CMD [ "port", "name", "hash" ]