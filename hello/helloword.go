package main

import (
	"encoding/binary"
	"fmt"
	"hello/msg"
	"net"

	"github.com/golang/protobuf/proto"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		panic(err)
	}
	test := &msg.Hello{
		// 使用辅助函数设置域的值
		Name: *proto.String("Leaf"),
	}
	// 进行编码
	data, err := proto.Marshal(test)
	if err != nil {
		fmt.Println("marshaling error: ", err)
	}
	// len + id + data
	m := make([]byte, 4+len(data))
	binary.BigEndian.PutUint16(m, uint16(len(data)+2))
	copy(m[4:], data)
	// 发送消息
	conn.Write(m)
	buf := make([]byte, 32)
	// 接收消息
	n, err := conn.Read(buf)
	if err != nil {

	}

	recv := &msg.Hello{}
	err = proto.Unmarshal(buf[4:n], recv)
	if err != nil {

	}
	fmt.Println(recv.GetMsgId())
	fmt.Println(recv.GetName())
}
