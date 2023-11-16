FROM golang:1.16

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

COPY --from=0 /app/main /app/main

EXPOSE 8123

CMD ["/app/main"]