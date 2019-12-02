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

- In directory cd /home/sandra/go/src/github.com/proyecto/backend_go

  - export GOROOT=/usr/local/go
  - export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
  - go version
  - go env
  - export GOPATH=/home/sandra/go/
  - export PATH=$PATH:/home/sandra/go/bin
  - echo $GOPATH
  - echo $PATH

- First time:

  - go get -u github.com/kardianos/govendor
  - go get -u golang.org/x/crypto/...
  - go get -u github.com/pilu/fresh

  - govendor sync
  - govendor add +external
  - fresh

- Other times:

  - fresh

- If port 8090 is being used, kill process:

  - sudo netstat -ltnp |grep :8090
  - sudo kill -9 PID

- Initializing at:

  - http://localhost:8090/api/articles/

- If inotify_init: too many open files problem:

  - cat /proc/sys/fs/inotify/max_user_instances
  - sudo su
  - echo 256 > /proc/sys/fs/inotify/max_user_instances

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

### SWAGGER

In directory cd /home/sandra/go/src/github.com/proyecto/swagger

  - export GOROOT=/usr/local/go
  - export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
  - export GOPATH=/home/sandra/go/
  - export PATH=$PATH:/home/sandra/go/bin

- First time:

  - go get -u github.com/swaggo/swag/cmd/swag 
  - go get -u github.com/swaggo/gin-swagger
  - go get -u github.com/swaggo/gin-swagger/swaggerFiles

  - go run main.go

- Other times:

  - go run main.go
  - http://0.0.0.0:3004/swagger/index.html


### BACKEND DOCKERIZED

- Build container:
   - sudo docker build -t back_go .

- See images:
   - sudo docker images -a

- Run container:
  - sudo docker run -d -p 8090:8090 --name back_go golang

- See container:

  - sudo docker ps -a

















