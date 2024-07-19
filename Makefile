all: proto

go:
	go build ./...

proto:
	protoc --proto_path=. --go_out=./generated proto/message.proto

clean:
	rm *.exe *.pb.go
