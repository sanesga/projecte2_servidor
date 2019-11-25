# projecte2_servidor

## Frontend

- Angular 

### Run Frontend Angular

- cd frontend_angular
- sudo npm install
- ng serve --host 0.0.0.0 --port 8081 --disableHostCheck true
- http://localhost:8081/

## Backend 

- Go 
- MySQL

### Run Backend Go

In directory /home/sandra/go

- export GOROOT=/usr/local/go
- export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
- go version
- go env
- export GOPATH=/home/sandra/go/
- export PATH=$PATH:/sandra/go/bin
- echo $GOPATH
- echo $PATH

- go get -u github.com/kardianos/govendor
- go get -u golang.org/x/crypto/...
- go get -u github.com/pilu/fresh

In directory /home/sandra/go/src/github.com/backend_go

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














