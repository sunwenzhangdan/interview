/*
 * @Date: 2021-10-13 17:44:12
 * @LastEditors: seven sun
 * @LastEditTime: 2021-10-14 08:42:31
 * @FilePath: /code/interview/gobase/functional/functional.go
 */

package functional

import "fmt"

func funcInner(say string) {
	c := say+"world"
	innerfunc := func(say string) {
		fmt.Println(say)
	}
	innerfunc(c)
}
