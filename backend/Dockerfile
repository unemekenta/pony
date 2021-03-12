FROM golang:1.16.1-alpine3.13

WORKDIR /go/src/app

ENV GO111MODULE=on

RUN apk add --no-cache \
    alpine-sdk \
    git \
    && go get github.com/pilu/fresh \
    && go get github.com/gorilla/handlers \
    && go get github.com/gorilla/mux \
    && go get github.com/jinzhu/gorm \
    && go get github.com/joho/godotenv \
    && go get bitbucket.org/liamstask/goose/cmd/goose \
    && go get github.com/lib/pq


EXPOSE 8080

CMD ["/bin/ash"]