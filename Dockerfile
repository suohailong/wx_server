FROM golang:1.8

WORKDIR /go/src/
COPY . .

RUN ls

RUN echo $GOPATH

RUN go get -d -v ./...

WORKDIR /go/src/wx_app/

