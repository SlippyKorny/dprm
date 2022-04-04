build:
	go build -o bin/dprm cmd/cli/main.go

install:
	go build -o $(GOPATH)/bin/dprm cmd/cli/main.go
	# TODO: installation for dprm-gui