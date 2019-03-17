FROM golang:1.8

WORKDIR /go/src/github.com/brogrammers-united/cli
COPY . .

RUN go get -d -v ./...
RUN go install

CMD ["go", "test"]