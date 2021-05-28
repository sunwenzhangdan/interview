/*
 * @Date: 2021-05-08 14:08:10
 * @LastEditors: seven sun
 * @LastEditTime: 2021-05-08 17:40:03
 * @FilePath: /面试题/gobase/slice/slice.go
 */
package slice

import "fmt"

/*
slice可以根据数组来创建，容量为从截取的首地址到结尾
*/
func GetSliceAddr() {
	var array [10]int
	var slicec = array[2:4]
	fmt.Println(len(slicec))
	fmt.Println(cap(slicec))
	fmt.Println(&array[5] == &slicec[0])
}

/*

 */
func GetDefaultSlice() {
	var array []int
	fmt.Println(len(array))
	fmt.Println(cap(array))
	appendx(array, 3)
	fmt.Println(len(array))
	fmt.Println(cap(array))
}

func appendx(slicex []int, element int) {
	slicex = append(slicex, 1, 2)
	slicex = append(slicex, element)
	fmt.Println("slicex", slicex)
}

func NewOrMake() {
	var user *User
	fmt.Println(user)
	user = new(User)
	fmt.Println(user)
	fmt.Println(&user)
}

type User struct {
	name string
	age  int
}
