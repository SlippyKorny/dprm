build:
	# cmd
	go build -o bin/dprm.exe cmd/cli/main.go

	# windows ui
	# rsrc -manifest cmd/dprmui/dprmui.manifest -o cmd/dprmui/rsrc.syso
	cp cmd/dprmui/dprmui.manifest bin/dprmui.exe.manifest
	# go build -ldflags="-H windowsgui" -o bin/dprmui.exe cmd/dprmui/*.go
	go build -o bin/dprmui.exe cmd/dprmui/*.go

install:
	go build -o $(GOPATH)/bin/dprm cmd/cli/main.go