package main

import (
	"bytes"
	"encoding/binary"
	"testing"
	"unsafe"
)

// 非安全指针修改[]byte
func Test_byte_unsafe_pointer(t *testing.T) {
	bytes := []byte("i is byte")
	str := (*string)(unsafe.Pointer(&bytes))
	bytes[0] = 'y'
	t.Log(*str)
}

type ResponseHeader struct {
	Length uint32 // 包长 4个字节
	Cmd    uint32 // 协议号 4个字节
}

func TestByte1(t *testing.T) {
	sendBuf := bytes.NewBuffer(make([]byte, 8))

	var headerLen uint32
	headerLen = 1111239090
	_ = binary.Write(sendBuf, binary.LittleEndian, headerLen)
	var protoNo int32
	protoNo = 222000
	_ = binary.Write(sendBuf, binary.LittleEndian, protoNo)
	t.Log(len(sendBuf.Bytes()))

	var length uint32 // 包长 4个字节
	b1 := bytes.NewBuffer(sendBuf.Bytes()[0:4])
	b2 := bytes.NewBuffer(sendBuf.Bytes()[4:])
	err := binary.Read(b1, binary.LittleEndian, &length)
	if err != nil {
		t.Log(err)
		return
	}

	var proton uint32 // 包长 4个字节
	err = binary.Read(b2, binary.LittleEndian, &proton)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(length)
	t.Log(proton)
}
