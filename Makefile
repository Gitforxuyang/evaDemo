
proto:
	mkdir ./proto/evaDemo || true
	protoc --eva_out=plugins=all:./proto/evaDemo -I=./proto evaDemo.proto
	protoc --go_out=plugins=grpc:./proto/evaDemo -I=./proto evaDemo.proto

swagger:
	protoc --swagger_out=plugins=grpc:./proto/evaDemo -I=./proto evaDemo.proto

.PHONY: proto

build:
	GOOS=linux GOARCH=amd64 go build -o evaDemo main.go

docker:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o evaDemo main.go
	docker build -t eva-demo:v1 .