APP      := cmd/app/main.go
MIGRATE  := cmd/migrate/main.go
SEED     := cmd/seed/main.go
ARCH     := amd64
OUTPUT   := build/monitor

run:
	CGO_ENABLED=1 go run $(APP)

migrate:
	CGO_ENABLED=0 go run $(MIGRATE)

seed:
	CGO_ENABLED=0 go run $(SEED)

build-linux:
	CGO_ENABLED=1 GOOS=linux GOARCH=$(ARCH) go build -o $(OUTPUT) $(APP)

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=$(ARCH) go build -o $(OUTPUT).exe $(APP)
