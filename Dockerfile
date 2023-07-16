FROM golang:1.20

WORKDIR /usr/src/app

COPY . .

RUN go mod download
RUN go build -o employees ./cmd/app/main.go

CMD ["./employees"]
