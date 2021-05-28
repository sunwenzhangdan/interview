/*
 * @Date: 2021-05-10 16:35:16
 * @LastEditors: seven sun 
 * @LastEditTime: 2021-05-10 16:43:56
 * @FilePath: /面试题/gobase/range/range_test.go
 */
package rangesun

import (
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {
	slicec := []int{1, 2, 3, 5, 7, 8, 8, 45}
	for i, _ := range slicec {
		fmt.Println(i, slicec[i])
	}
}
