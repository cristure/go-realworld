
FROM golang:1.19

ADD . /go/src/go-realworld
WORKDIR /go/src/go-realworld
RUN go get github.com/go-realworld
RUN go install
ENTRYPOINT ["/go/bin/go-realworld"]
