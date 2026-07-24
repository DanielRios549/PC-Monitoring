run:
	CGO_ENABLED=1 go run main.go

build:
	CGO_ENABLED=1 go build -o build/monitor

build-linux:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o build/monitor

build-windows:
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o build/monitor.exe
