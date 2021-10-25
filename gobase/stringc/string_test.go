/*
 * @Date: 2021-10-20 12:59:10
 * @LastEditors: seven sun
 * @LastEditTime: 2021-10-21 14:50:58
 * @FilePath: /interview/gobase/stringc/string_test.go
 */
package stringc

import (
	"fmt"
	"testing"
)

func TestStringc(t *testing.T) {
	b := "heelo中国"
	for i, c := range b {

		fmt.Println(i)
		fmt.Println(c)
	}
	t.Log("c")
}
