FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /go-crud-app

EXPOSE 8000

CMD ["./go-crud-app"]
