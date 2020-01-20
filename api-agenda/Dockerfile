FROM golang

ADD ./src /go/src/api-agenda

RUN go get github.com/gorilla/mux
RUN go get github.com/lib/pq

RUN go install api-agenda

ENTRYPOINT /go/bin/api-agenda

EXPOSE 8080
