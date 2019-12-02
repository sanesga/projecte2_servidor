FROM golang:latest
COPY . /go/src/github.com/proyecto/backend_go
ENV GOPATH /go
ENV GOROOT /usr/local/go
ENV export GOPATH=/home/sandra/go/
ENV export PATH=$PATH:/home/sandra/go/bin

RUN go get -u github.com/kardianos/govendor; \
    go get -u github.com/pilu/fresh; \
    go get -u golang.org/x/crypto/... ;
ADD ./vendor /go/src/github.com/proyecto/backend_go/vendor

WORKDIR /go/src/github.com/proyecto/backend_go
RUN govendor sync;
RUN govendor add +external
EXPOSE 8090
ENTRYPOINT ["fresh"]