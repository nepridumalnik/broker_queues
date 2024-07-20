all: go

go: protoc
	go build ./...

protoc:
	protoc --proto_path=. --go_out=./generated proto/message.proto

clean:
	rm *.exe *.pb.go
