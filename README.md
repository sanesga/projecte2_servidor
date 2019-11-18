# projecte2_servidor

## Frontend

- Angular 

## Backend 

- Go 
- MySQL

### For running

In directory /home/sandra/go

- export GOROOT=/usr/local/go
- export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
- go version
- go env
- export GOPATH=/home/sandra/go/
- export PATH=$PATH:/sandra/go/bin
- echo $GOPATH
- echo $PATH

cd src/github.com/backend_go

- go get -u github.com/kardianos/govendor
- go get -u golang.org/x/crypto/...
- go get -u github.com/pilu/fresh

- govendor sync 
- govendor add +external
- fresh

kill PID at port 8090

- sudo netstat -ltnp |grep :8090
- sudo kill -9 PID

initializing at 

- http://localhost:8090/api/articles/

###  POSTMAN

- Get a Token:

POST http://localhost:8090/api/users/login
{
  "user":{
    "username": "sandra",
    "email": "sandra@gmail.com",
    "password": "12345678"
  }
}

- See data with authorization Token

GET http://localhost:8090/api/articles
GET http://localhost:8090/api/books
GET http://localhost:8090/api/books/slug














