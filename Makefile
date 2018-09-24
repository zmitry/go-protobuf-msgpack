dep:
	brew install protobuf
	go get -u github.com/golang/protobuf/{proto,protoc-gen-go}

protoc:
	@mkdir -p proto
	protoc --go_out=proto *.proto

run:
	go run server.go

test:
	go test .

bench:
	go test -bench . -benchmem
