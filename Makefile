
build:
	GOPROXY=direct CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/main ./main.go

start:
	./bin/main
