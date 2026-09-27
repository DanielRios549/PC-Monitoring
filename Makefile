FILE     := cmd/app/main.go
ARCH     := amd64
OUTPUT   := build/monitor

run:
	CGO_ENABLED=1 go run $(FILE)

build-linux:
	CGO_ENABLED=1 GOOS=linux GOARCH=$(ARCH) go build -o $(OUTPUT) $(FILE)

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=$(ARCH) go build -o $(OUTPUT).exe $(FILE)
