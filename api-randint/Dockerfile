FROM golang

ADD ./src /go/src/api-randint

RUN go get github.com/gorilla/mux

RUN go install api-randint

ENTRYPOINT /go/bin/api-randint

EXPOSE 5000
