COMPILER := x86_64-w64-mingw32-gcc
HEADERS  := "-I/usr/local/include"
LIBS     := "-L/usr/local/lib"
ARCH     := amd64
OUTPUT   := build/monitor

run:
	CGO_ENABLED=1 go run main.go

run-debug:
	CGO_ENABLED=1 ZE_ENABLE_LOADER_DEBUG_TRACE=1 go run main.go

build-linux:
	CGO_ENABLED=1 GOOS=linux GOARCH=$(ARCH) go build -o $(OUTPUT) main.go

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=$(ARCH) CC=$(COMPILER) CGO_CFLAGS=$(HEADERS) CGO_LDFLAGS=$(LIBS) go build -o $(OUTPUT).exe main.go
