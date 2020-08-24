gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb --go_opt=paths=source_relative

clean:
	rm pb/*.go

server:
	go run cmd/server/main.go -port 8088

client:
	go run cmd/client/main.go -address 0.0.0.0:8088

test:
	go test -cover -race ./...