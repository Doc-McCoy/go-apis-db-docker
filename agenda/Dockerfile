FROM golang

ADD . /src

RUN go get github.com/lib/pq
RUN go install main

ENTRYPOINT /src

EXPOSE 8080
