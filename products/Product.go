// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package products

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Product struct {
	_tab flatbuffers.Table
}

func GetRootAsProduct(buf []byte, offset flatbuffers.UOffsetT) *Product {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Product{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsProduct(buf []byte, offset flatbuffers.UOffsetT) *Product {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Product{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Product) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Product) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Product) Id() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Product) MutateId(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func (rcv *Product) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Product) Description() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Product) Price() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Product) MutatePrice(n int32) bool {
	return rcv._tab.MutateInt32Slot(10, n)
}

func (rcv *Product) Colors(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *Product) ColorsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func ProductStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func ProductAddId(builder *flatbuffers.Builder, id uint64) {
	builder.PrependUint64Slot(0, id, 0)
}
func ProductAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(name), 0)
}
func ProductAddDescription(builder *flatbuffers.Builder, description flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(description), 0)
}
func ProductAddPrice(builder *flatbuffers.Builder, price int32) {
	builder.PrependInt32Slot(3, price, 0)
}
func ProductAddColors(builder *flatbuffers.Builder, colors flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(colors), 0)
}
func ProductStartColorsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ProductEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
