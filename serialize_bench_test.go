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
			Id:          191919191919,
			Name:        "FIFA World Cup Soccer Ball",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
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
			Id:          191919191919,
			Name:        "FIFA World Cup Soccer Ball",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
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
			Id:          191919191919,
			Name:        "FIFA World Cup Soccer Ball",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
			Price:       18900,
			Colors:      []string{"red", "yello", "blue"},
		}

		buf, _ := json.Marshal(test)

		newTest := &product{}
		_ = json.Unmarshal(buf, newTest)
	}
}
