FROM golang:1.17

WORKDIR /go/src/blackjack
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o app .

EXPOSE 8080

ENTRYPOINT ["/go/src/blackjack/app"]

