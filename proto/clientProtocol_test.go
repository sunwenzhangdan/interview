/*
 * @Date: 2021-08-15 15:45:08
 * @LastEditors: seven sun
 * @LastEditTime: 2021-08-15 17:11:49
 * @FilePath: /面试题/proto/clientProtocol_test.go
 */
package proto

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"testing"
)

func Encode(message string) ([]byte, error) {
	length := int16(len(message))
	nb := new(bytes.Buffer)
	err := binary.Write(nb, binary.LittleEndian, length)
	if err != nil {
		fmt.Println(err)
	}
	err = binary.Write(nb, binary.LittleEndian, []byte(message))
	if err != nil {
		fmt.Println(err)
	}
	return nb.Bytes(), nil
}

func TestClient(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:8787")
	if err != nil {
		fmt.Println("net dial error", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 30; i++ {
		msg := `hello world,hello tom`
		data, err := Encode(msg)
		if err != nil {
			fmt.Println(err)
		}
		conn.Write(data)
	}
}
