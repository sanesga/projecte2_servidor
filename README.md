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
