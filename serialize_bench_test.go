package main

import (
	"encoding/json"
	"testing"

	"github.com/golang/protobuf/proto"
	_proto "github.com/uqichi/go-protobuf-msgpack/proto"
	"github.com/vmihailenco/msgpack"
)

func BenchmarkProtobuf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test := &_proto.Product{
			Id:          99,
			Name:        "Soccer Ball",
			Description: "This is a golden ball!!",
			Price:       18900,
			Colors:      []string{"red", "yello", "blue"},
		}

		buf, _ := proto.Marshal(test)

		newTest := &_proto.Product{}
		_ = proto.Unmarshal(buf, newTest)
	}
}

func BenchmarkMsgpack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test := &product{
			Id:          99,
			Name:        "Soccer Ball",
			Description: "This is a golden ball!!",
			Price:       18900,
			Colors:      []string{"red", "yello", "blue"},
		}

		buf, _ := msgpack.Marshal(test)

		newTest := &product{}
		_ = msgpack.Unmarshal(buf, newTest)
	}
}

func BenchmarkJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test := &product{
			Id:          99,
			Name:        "Soccer Ball",
			Description: "This is a golden ball!!",
			Price:       18900,
			Colors:      []string{"red", "yello", "blue"},
		}

		buf, _ := json.Marshal(test)

		newTest := &product{}
		_ = json.Unmarshal(buf, newTest)
	}
}
