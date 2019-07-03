FROM golang:latest

WORKDIR $GOPATH/src/books

COPY . .

RUN go get -d -v ./...

RUN go build

RUN go install -v ./...

EXPOSE 8080

CMD ["books"]