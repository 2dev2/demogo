

#if you want to run client and server seperate, then use client and server
#if client and server on same server address then use app

export VERSION="1.16.0"
export ARCH="amd64"

.PHONY: init
init:
	curl -sL https://raw.githubusercontent.com/nvm-sh/nvm/v0.35.0/install.sh -o install_nvm.sh
	bash install_nvm.sh
	source ~/.bash_profile
	nvm install --lts
	cd web && npm run install
	curl -O -L "https://golang.org/dl/go${VERSION}.linux-${ARCH}.tar.gz"
	tar -xf "go${VERSION}.linux-${ARCH}.tar.gz"
	mv -v go /usr/local
	export GOPATH=$HOME/go
	export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
	source ~/.bash_profile
	go version
	go mod tidy

.PHONY: client
client:
	cd web && npm run start

.PHONY: server
server:
	go run ./bin/main.go

.PHONY: app
app:
	cd web && npm run build
	go mod tidy
	go run ./bin/main.go


