/*
 * @Date: 2021-05-12 07:23:02
 * @LastEditors: seven sun
 * @LastEditTime: 2021-05-12 08:09:08
 * @FilePath: /面试题/gobase/datastruct/hashmap/hashmap_test.go
 */

package hashmap

import (
	"math"
	"strings"
)

type Node struct {
	key   string
	value string
	hash  int
	next  *Node
}

const CAPACITY = 16

//创建存放数据的地方
var table = make([]*Node, CAPACITY)
var size int

func isEmpty() bool {
	if size == 0 {
		return true
	} else {
		return false
	}
}

//传入一个key生成hashcode
func hashcode(key string) int {
	var num = 0
	var len = len(key)
	for i := 0; i < len; i++ {
		num += int(key[i])
	}
	var avg = num * int((math.Pow(5.0, 5.0) - 1)) / 2
	var numeric = avg - int(math.Floor(float64(avg)))
	return int(math.Floor(float64(numeric * CAPACITY)))
}

/*
放入一个数据
*/
func put(key, value string) {
	var hash = hashcode(key)
	var newNode *Node = new(Node)
	newNode.key = key
	newNode.value = value
	newNode.hash = hash
	newNode.next = nil
	var node = table[hash]
	//hahs是一个数字，是在数组的索引
	//去当前结构体数组查看是否有个这个元素，如果没有
	//但是这个查的是一个链表,如果没查到，就直接加入
	for {
		if node == nil {
			break
		}
		if strings.Compare(node.key, key) == 0 {
			node.value = value
			return
		}
		node = node.next
	}

	newNode.next = table[hash]
	table[hash] = newNode
	size++

}

func get(key string) string {
	if key == "" {
		return ""
	}
	var hash = hashcode(key)
	var node = table[hash]
	for {
		if node == nil {
			break
		}
		if strings.Compare(node.key, key) == 0 {
			return node.key
		}
		node = node.next
	}
	return ""
}
