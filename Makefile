ARCH     := amd64
OUTPUT   := build/monitor

run:
	CGO_ENABLED=1 go run main.go

build-linux:
	CGO_ENABLED=1 GOOS=linux GOARCH=$(ARCH) go build -o $(OUTPUT) main.go

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=$(ARCH) go build -o $(OUTPUT).exe main.go
