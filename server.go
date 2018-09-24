package main

import (
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"
	_proto "github.com/uqichi/go-protobuf-msgpack/proto"
	"github.com/vmihailenco/msgpack"
)

func main() {
	http.HandleFunc("/pb", handlerProtobuf)
	http.HandleFunc("/mp", handlerMsgpack)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handlerProtobuf(w http.ResponseWriter, r *http.Request) {
	p := &_proto.Product{
		Id:          99,
		Name:        "Soccer Ball",
		Description: "This is a golden ball!!",
		Price:       18900,
		Colors:      []string{"red", "yello", "blue"},
	}

	res, err := proto.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/protobuf")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

type MsgpackProduct struct {
	Id          int64
	Name        string
	Description string
	Price       int32
	Colors      []string
}

func handlerMsgpack(w http.ResponseWriter, r *http.Request) {
	p := &MsgpackProduct{
		Id:          99,
		Name:        "Soccer Ball",
		Description: "This is a golden ball!!",
		Price:       18900,
		Colors:      []string{"red", "yello", "blue"},
	}

	res, err := msgpack.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/x-msgpack")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
