/*
 * @Date: 2021-05-08 18:22:01
 * @LastEditors: seven sun
 * @LastEditTime: 2021-10-22 14:54:15
 * @FilePath: /code/interview/gobase/channel/channel_test.go
 */
package channelsun

import (
	"fmt"
	"testing"
	"time"
)

// func TestCloseChannel(t *testing.T) {
// 	var ch chan struct{} = make(chan struct{})
// 	go func() {
// 		for {
// 			fmt.Println(<-ch)
// 		}
// 	}()
// 	go func() {
// 		for {
// 			fmt.Println(<-ch)
// 		}
// 	}()
// 	ch <- struct{}{}
// 	close(ch)
// 	time.Sleep(20 * time.Second)
// }

func TestSelectNil(t *testing.T) {
	a := make(chan int)
	b := make(chan int)
	var c int
	go func() {
		for {
			select {
			case a <- 1:
				c = <-a
				fmt.Println(c)
				a = nil
				b = make(chan int)
			case b <- 2:
				c = <-b
				fmt.Println(c)
				b = nil
				a = make(chan int)
			}
		}
	}()
	go func() {
		for {
			fmt.Println("cccccc", c)
		}
	}()
	time.Sleep(time.Second * 100)
}
