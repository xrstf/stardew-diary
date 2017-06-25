default: build

build: fix
	go build -v .

fix: *.go
	goimports -l -w .
	gofmt -l -w .

data:
	cd sdv/data/src && go run main.go
