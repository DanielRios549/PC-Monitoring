run:
	go run main.go

build:
	go build -o build/monitor

build-linux:
	GOOS=linux GOARCH=amd64 go build -o build/monitor

build-windows:
	GOOS=windows GOARCH=amd64 go build -o build/monitor.exe
