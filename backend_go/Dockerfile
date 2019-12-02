FROM golang:latest
WORKDIR /go/src/github.com/proyecto/backend_go
RUN go get -u github.com/kardianos/govendor; \
    go get -u github.com/pilu/fresh; \
    go get -u golang.org/x/crypto/... ;
COPY ./vendor ./vendor
RUN /usr/bin/sqlite3 /db/test.db
COPY . .
EXPOSE 8090
CMD ["fresh"]



  