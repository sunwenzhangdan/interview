/*
 * @Date: 2021-08-15 15:51:16
 * @LastEditors: seven sun
 * @LastEditTime: 2021-08-15 17:11:26
 * @FilePath: /面试题/proto/serverProtocol_test.go
 */
package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"testing"
)

func Decode(reader *bufio.Reader) (string, error) {
	lengthByte, _ := reader.Peek(2) // 读取前两个字节
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int16
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	// 如果所读的数据小于一个包的数据就先不读了
	if int16(reader.Buffered()) < length+2 {
		return "", err
	}
	realData := make([]byte, int(2+length))
	_, err = reader.Read(realData)
	if err != nil {
		return "", err
	}
	return string(realData[2:]), nil
}

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := Decode(reader)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(msg)
	}
}

func TestServer(t *testing.T) {
	listen, err := net.Listen("tcp", "0.0.0.0:8787")
	if err != nil {
		fmt.Println(err)
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		go process(conn)
	}
}
