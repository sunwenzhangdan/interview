/*
 * @Date: 2021-08-13 13:46:50
 * @LastEditors: seven sun
 * @LastEditTime: 2021-08-15 10:30:02
 * @FilePath: /面试题/leetcode/twosum_test.go
 */
package leetcode

import (
	"net"
	"testing"
)

func TestTwoSum(t *testing.T) {
	t.Log("hello")
	t.Log(Sum([]int{1, 3, 4, 6, 7, 9}, 10))
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = listener.Close()
	}()
	t.Logf("bound to %q", listener.Addr())
}
