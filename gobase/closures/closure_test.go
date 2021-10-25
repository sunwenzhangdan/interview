/*
 * @Date: 2021-08-11 20:29:05
 * @LastEditors: seven sun
 * @LastEditTime: 2021-08-12 07:46:53
 * @FilePath: /面试题/gobase/closures/closure_test.go
 */
package closures

import (
	"fmt"
	"net/http"
	"testing"
)

type Gen func() int

// 斐波那契
func makeGen() Gen {
	i := 0
	j := 1
	return func() int {
		c := i + j
		i, j = j, i+j
		return c
	}
}

func createHttp() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":3003", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>hello</h1>")
}

func TestGen(t *testing.T) {
	c := makeGen()
	for i := 0; i < 10; i++ {
		fmt.Println(c())
	}
	createHttp()
}

//创建中间件
