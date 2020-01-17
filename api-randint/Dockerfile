FROM golang

ADD . /src

RUN go get github.com/gorilla/mux
RUN go install main

ENTRYPOINT /src

EXPOSE 8081
