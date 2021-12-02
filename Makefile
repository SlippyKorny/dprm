build:
	go build -o bin/dprm cmd/cli/main.go
	go build -o bin/dprm-gui cmd/gui/main.go

install:
	go build -o $(GOPATH)/bin/dprm cmd/cli/main.go
	# TODO: installation for dprm-gui