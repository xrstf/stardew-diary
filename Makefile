default: build

build: fix
	go build -v .

fix: *.go
	goimports -l -w .
	gofmt -l -w .

data:
	cd sdv/data/src && go run main.go

release:
	GOOS=windows GOARCH=386 go build -v .
	mkdir -p stardew-diary
	mv stardew-diary.exe stardew-diary/
	cp "Command Prompt.lnk" stardew-diary/
	unix2dos -n README.md stardew-diary/README.txt
	cp D:\bin\fossil.exe stardew-diary/
	zip -r stardew-diary.zip stardew-diary/
