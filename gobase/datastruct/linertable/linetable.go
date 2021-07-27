/*
 * @Date: 2021-05-11 07:32:19
 * @LastEditors: seven sun
 * @LastEditTime: 2021-07-19 11:19:01
 * @FilePath: /面试题/gobase/datastruct/linertable/linetable.go
 */
package linertable

import (
	"fmt"
)

/*
对线性数据的迭代
*/
func seekscores() {
	scores := []int{1, 2, 4, 6, 3, 5, 9, 8, 7}
	lenth := len(scores)
	for i := 0; i < lenth; i++ {
		fmt.Printf("%d", scores[i])
	}
}

/*
找出最大值
*/
func findMax(slicex []int) int {
	for i := 0; i < len(slicex); i++ {
		if slicex[i] > slicex[i+1] {
			slicex[i], slicex[i+1] = slicex[i+1], slicex[i]
		}
	}
	return slicex[len(slicex)-1]
}

/*
通过以上方法我们就得到了冒泡算法,第一次找出最大，第二次找到次大
*/

func MaoPao(slicex []int) {
	for j := len(slicex); j > 0; j-- {
		for i := 0; i < j; i++ {
			if slicex[i] > slicex[i+1] {
				slicex[i], slicex[i+1] = slicex[i+1], slicex[i]
			}
		}
	}
}

/*
冒泡会频繁交换，我们下面采用一种方法，先比较，找打最大值或者最小值直接交换位置
需要借用一块多余的内存保存目前最小值的位置
*/
//找出最小值的索引位置
func min(slicex []int) {
	minindex := 0
	// 经过这一轮得出的是最小值的位置
	for i := 1; i < len(slicex); i++ {
		if slicex[i] < slicex[minindex] {
			minindex = i
		}
	}
}

/*
选择排序
*/

func Xuanze(slicex []int) {
	for j := 0; j < len(slicex); j++ {
		// 默认大小为当前
		minindex := j
		for i := 1; i < len(slicex); i++ {
			if slicex[i] < slicex[minindex] {
				minindex = i
			}
		}
		slicex[j], slicex[minindex] = slicex[minindex], slicex[j]
	}
}

/*
在数组后面添加一个数字
*/
func appendx(array []int, index int) {
	leg := len(array)
	temparray := make([]int, 0, leg+1)
	for i := 0; i < leg; i++ {
		temparray[i] = array[i]
	}
	temparray[leg] = index
}

/*
在数组中间插入一个
*/
func insertIndex(slicex []int, index int, num int) {
}

/*
插入排序
*/
func insertSort(slicex []int) {
	for j := 1; j < len(slicex); j++ {
		// j是代表第j个数插入到前面的有序队列里面
		insertnum := slicex[j]
		i := j - 1
		for ; i >= 0; i-- {
			if slicex[i] > insertnum {
				slicex[i+1] = slicex[i]
			} else {
				break
			}
		}
		slicex[i+1] = insertnum
	}
}

/*
翻转数组
*/
func reversArray(slicex []int) {
	i := 0
	j := len(slicex) - 1
	for i <= j {
		slicex[i], slicex[j] = slicex[j], slicex[i]
		i++
		j--
	}
}

/*
二分查找
*/

func dividesearch(slicex []int, target int) int {
	// 判断这个合理不合理，主要靠[1,1] 看这个空间是有效的吗
	start := 0
	end := len(slicex) - 1
	for start <= end {
		mid := (start + end) / 2
		if slicex[mid] == target {
			return mid
		} else if slicex[mid] < target {
			start = mid + 1
		} else if slicex[mid] > target {
			end = mid - 1
		} else {
			return -1
		}
	}
	return -1
}

func Guibing(array []int) {
}

// 什么叫递归，就是把一个大问题，拆成子问题，子问题继续拆，拆到直到无法继续拆，跳出递归
// func guibingSort(array []int, l, r int) {
// 	if l >= r {
// 		return
// 	}
// 	mid := (l + r) / 2
// 	guibingSort(array, l, mid)
// 	guibingSort(array, mid+1, r)
// 	MergeGuibing(array, l, mid, r)
// }

// func MergeGuibing(array []int, l, mid, r int) {
// 	temparray := make([]int, r-l+1)
// 	// 切片拷贝
// 	copy(temparray, array[l:r+1])
// 	fmt.Println(temparray)
// 	i := l
// 	j := mid + 1
// }

// func Guibi1(array []int) []int {
// 	return Guibisort1(array)
// }

// func Guibisort1(array []int) []int {
// 	fmt.Println(array)
// 	if len(array) < 2 {
// 		return array
// 	}
// 	mid := len(array) / 2
// 	lslice := Guibisort1(array[:mid])
// 	rslice := Guibisort1(array[mid+1:])
// 	sorted := GuibiMerge(lslice, rslice)
// 	return sorted
// }

// func GuibiMerge(lslice, rslice []int) []int {
// 	i := 0
// 	j := 0
// 	temp := make([]int, 1)
// 	for i <= len(lslice)-1 && j <= len(rslice)-1 {
// 		if lslice[i] < rslice[j] {
// 			temp = append(temp, lslice[i])
// 			i++
// 		} else {
// 			temp = append(temp, rslice[j])
// 			j++
// 		}
// 		if i < len(lslice)-1 {
// 			temp = append(temp, lslice[i])
// 			i++

// 		}
// 		if j < len(rslice)-1 {
// 			temp = append(temp, rslice[j])
// 			j++
// 		}
// 	}
// 	return temp
// }

// 采用返回数据
func MergeSort(arr []int) []int {
	fmt.Println(arr)
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2           // 一分为二
	left := MergeSort(arr[0:mid]) // 将数组的前半部分递归
	right := MergeSort(arr[mid:]) // 将数组的后半部分递归
	result := merge(left, right)  // 分解的传递给合并函数
	return result
}

// 将分解的小数组合并成一个大数组
func merge(left, right []int) []int {
	result := make([]int, 0)          // 申请一个临时空间
	m, n := 0, 0                      // 申请二个指针
	l, r := len(left)-1, len(right)-1 // 申请二个数组的边界指针
	for m <= l && n <= r {
		if left[m] < right[n] { // 小于则加入临时数组
			result = append(result, left[m])
			m++
		} else {
			result = append(result, right[n])
			n++
		}
	}
	if m <= l { // 将未移动完的数据直接加入临时数组
		result = append(result, left[m:]...)
	}
	if n <= r { // 将未移动完的数据直接加入临时数组
		result = append(result, right[n:]...)
	}

	return result
}

func QuickSort(slicex []int, startx, endx int) {
	if startx >= endx {
		return
	}
	mid := Pratition(slicex, startx, endx)
	QuickSort(slicex, startx, mid-1)
	QuickSort(slicex, mid+1, endx)
}

func Pratition(slicex []int, start, end int) int {
	i, j, k := start, end-1, slicex[end]
	for i != j {
		for i < j && slicex[i] < k {
			i++
		}
		for i < j && slicex[j] >= k {
			j--
		}
		if i < j {
			slicex[i], slicex[j] = slicex[j], slicex[i]
			i++
			j--
		}

	}
	slicex[start], slicex[i] = slicex[i], sicex[start]
	return i
}
