/*
 * @Date: 2021-05-12 10:08:07
 * @LastEditors: seven sun
 * @LastEditTime: 2021-11-25 16:49:03
 * @FilePath: /interview/gobase/context/context_test.go
 */
package contextsun

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	// ceshiAdd()
	// stopGo()
	// byConetxt()
	// b:=returf()
	// b()
	testClosechannel()
}

func ceshiAdd() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("i am one")
		time.Sleep(time.Second * 2)
		wg.Done()
	}()
	go func() {
		fmt.Println("i am two")
		time.Sleep(time.Second * 2)
		wg.Done()
	}()
	wg.Wait()
}

// 现在我们又有一个需求，我们想主动停下某个工作
func stopGo() {
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Printf("%v", cancel)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("收到信号")
				return
			default:
				fmt.Println("i am working")
			}
		}
	}(ctx)
	time.Sleep(time.Second * 1)
	cancel()
}

func byConetxt() {
	ctx := context.Background()
	ctx1 := context.WithValue(ctx, "k1", "v1")
	ctx2 := context.WithValue(ctx1, "k2", "v2")
	ctx3 := context.WithValue(ctx2, "k3", "v3")
	func1(ctx3)
}

func func1(ctx context.Context) {
	fmt.Println(ctx.Value("k4").(string))
}

type f1 func()

func sum(a, b int) int {
	return a + b
}

func returf() f1 {
	return func() {
		a := sum(1, 3)
		fmt.Println(a)
	}
}

func testClosechannel() {
	chan1 := make(chan struct{})
	go func() {
		for {
			select {
			case <-chan1:
				fmt.Println("11111")
				return
			default:
				fmt.Println("2222")
			}
		}
	}()
	go func() {
		for {
			select {
			case <-chan1:
				fmt.Println("33333")
				return
			default:
				fmt.Println("44444")
			}
		}
	}()
	time.Sleep(time.Second * 2)
	close(chan1)
	time.Sleep(time.Second * 2)
}

func hello(ctx context.Context) {
	
}
