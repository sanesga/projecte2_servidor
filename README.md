# LIBRARY PROJECT

Functions:

- Login and register
- Book's list and details
- Social login (github)
- Book comments
- Favourite books

## Frontend

- Angular 

## Backend Go

- Go
- SQLite
- Docker
- Redis

## Backend Swagger

- Go
- Swagger


### Run Frontend Angular

- cd frontend_angular
- sudo npm install
- ng serve --host 0.0.0.0 --port 8081 --disableHostCheck true
- http://localhost:8081/


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
  - go get -u github.com/go-redis/redis

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
  - echo 10000 > /proc/sys/fs/inotify/max_user_instances

### Postman

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

### Backend Dockerized

- Build image:
   - sudo docker build -t back_go .

- See images:
   - sudo docker images -a

- Build container:
  - sudo docker run -p 8090:8090 --name backend_go_container back_go

- See container:

  - sudo docker ps -a

- If inotify_init: too many open files problem:

  - cat /proc/sys/fs/inotify/max_user_instances
  - sudo su
  - echo 10000 > /proc/sys/fs/inotify/max_user_instances

- Delete old containers:

  - sudo docker rm containerName

### Docker compose

- sudo docker-compose up

- http://localhost:8081/

Para sobreescribir la imagen:

- sudo docker-compose up --build


### Run Backend Swagger 

Falta:
 - añadir localhost:8090/api/users -- get all users
 - comments
 - favorite

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


POST
http://localhost:8090/
{
	"key": "sandra",
	"value": "esplugues"
}
GET
http://localhost:8090/api/redis/sandra
(Enviem la key)