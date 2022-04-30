package main

import (
	"test/products"
	"testing"

	"capnproto.org/go/capnp/v3"
	json "github.com/goccy/go-json"
	flatbuffers "github.com/google/flatbuffers/go"

	"test/foo"

	"github.com/golang/protobuf/proto"
	_proto "github.com/uqichi/go-protobuf-msgpack/proto"
	"github.com/vmihailenco/msgpack"
)

func BenchmarkProtobuf(b *testing.B) {
	test := &_proto.Product{
		Id:          191919191919,
		Name:        "FIFA World Cup Soccer Ball",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
		Price:       18900,
		Colors:      []string{"red", "yello", "blue"},
	}

	buf, _ := proto.Marshal(test)

	newTest := &_proto.Product{}
	for i := 0; i < b.N; i++ {
		newTest = &_proto.Product{}
		_ = proto.Unmarshal(buf, newTest)
	}
}

func BenchmarkMsgpack(b *testing.B) {
	test := &product{
		Id:          191919191919,
		Name:        "FIFA World Cup Soccer Ball",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
		Price:       18900,
		Colors:      []string{"red", "yello", "blue"},
	}

	buf, _ := msgpack.Marshal(test)

	newTest := &product{}
	for i := 0; i < b.N; i++ {

		_ = msgpack.Unmarshal(buf, newTest)
	}
}

func BenchmarkJSON(b *testing.B) {
	test := &product{
		Id:          191919191919,
		Name:        "FIFA World Cup Soccer Ball",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
		Price:       18900,
		Colors:      []string{"red", "yello", "blue"},
	}

	buf, _ := json.Marshal(test)

	newTest := &product{}
	for i := 0; i < b.N; i++ {
		newTest = &product{}
		_ = json.Unmarshal(buf, newTest)
	}
}

func BenchmarkFlatbuffer(b *testing.B) {
	builder := flatbuffers.NewBuilder(0)
	s1 := builder.CreateString("FIFA World Cup Soccer Ball")
	s2 := builder.CreateString("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
	products.ProductStart(builder)
	products.ProductAddId(builder, 191919191919)
	products.ProductAddName(builder, s1)
	products.ProductAddDescription(builder, s2)
	products.ProductAddPrice(builder, 18900)
	res := products.ProductEnd(builder)
	builder.Finish(res)

	buf := builder.FinishedBytes()

	var d string
	for i := 0; i < b.N; i++ {
		p := products.GetRootAsProduct(buf, 0)
		d = string(p.Description())
		t(d)
	}
	d = ""
	print(d)
}

func t(s string) {

}

func BenchmarkCapnproto(b *testing.B) {
	// Make a brand new empty message.  A Message allocates Cap'n Proto structs.
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		panic(err)
	}

	// Create a new Book struct.  Every message must have a root struct.
	product, err := foo.NewProduct(seg)

	if err != nil {
		panic(err)
	}
	product.SetId(191919191919)
	product.SetName("FIFA World Cup Soccer Ball")
	product.SetDescription("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
	product.SetPrice(18900)
	c, err := product.NewColors(3)
	if err != nil {
		panic("could not create colors array: " + err.Error())
	}
	c.Set(0, "red")
	c.Set(1, "yello")
	c.Set(2, "blue")
	product.SetColors(c)
	buf, err := msg.Marshal()
	if err != nil {
		panic(err)
	}
	var d string
	for i := 0; i < b.N; i++ {
		msg, err = capnp.Unmarshal(buf)
		res, _ := foo.ReadRootProduct(msg)
		d, _ := res.Description()
		t(d)
	}
	d = ""
	print(d)
}
