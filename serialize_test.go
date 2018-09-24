package main

import (
	"encoding/json"
	"testing"

	"github.com/golang/protobuf/proto"
	_proto "github.com/uqichi/go-protobuf-msgpack/proto"
	"github.com/vmihailenco/msgpack"
)

func TestProtobuf(t *testing.T) {
	test := &_proto.Product{
		Id:          99,
		Name:        "Soccer Ball",
		Description: "This is a golden ball!!",
		Price:       18900,
		Colors:      []string{"red", "yello", "blue"},
	}

	buf, err := proto.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}

	newTest := &_proto.Product{}
	err = proto.Unmarshal(buf, newTest)
	if err != nil {
		t.Fatal(err)
	}

	if test.Name != newTest.Name {
		t.Fatal("not matched")
	}
}

func TestMsgpack(t *testing.T) {
	test := &product{
		Id:          99,
		Name:        "Soccer Ball",
		Description: "This is a golden ball!!",
		Price:       18900,
		Colors:      []string{"red", "yello", "blue"},
	}

	buf, err := msgpack.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}

	newTest := &product{}
	err = msgpack.Unmarshal(buf, newTest)
	if err != nil {
		t.Fatal(err)
	}

	if test.Name != newTest.Name {
		t.Fatal("not matched")
	}
}

func TestJSON(t *testing.T) {
	test := &product{
		Id:          99,
		Name:        "Soccer Ball",
		Description: "This is a golden ball!!",
		Price:       18900,
		Colors:      []string{"red", "yello", "blue"},
	}

	buf, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}

	newTest := &product{}
	err = json.Unmarshal(buf, newTest)
	if err != nil {
		t.Fatal(err)
	}

	if test.Name != newTest.Name {
		t.Fatal("not matched")
	}
}
