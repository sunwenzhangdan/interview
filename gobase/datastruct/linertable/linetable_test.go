/*
 * @Date: 2021-05-11 16:54:03
 * @LastEditors: seven sun 
 * @LastEditTime: 2021-05-19 19:51:32
 * @FilePath: /面试题/gobase/datastruct/linertable/linetable_test.go
 */
package linertable

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	slicex := []int{1, 3, 5, 4, 11, 8, 6, 7, 10}
	insertSort(slicex)
	t.Log(slicex)
	reversArray(slicex)
	t.Log("hello")
	t.Log(slicex)
	fmt.Println(dividesearch([]int{1, 2, 3, 5, 7, 8, 9, 10}, 10))
	sci := []int{1, 5, 2, 4, 3, 7, 9, 10, 45, 595, 21, 455, 43, 0}
	QuickSort(sci, 0, 14)
	fmt.Println(sci)
}
