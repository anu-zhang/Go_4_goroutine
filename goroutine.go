package main

import (
	"sync"
	"fmt"
	"runtime"
)

// 第一次测试
//func main() {
//
//	for i := 0; i < 10; i++ {
//		go Add(1, 1)
//	}
//}
//
//func Add(x, y int) {
//	z := x + y
//	fmt.Println(z)
//
//}

// 第二个例子
//var counter int = 0
//
//func Count(lock *sync.Mutex) {
//	lock.Lock()
//	counter++
//	fmt.Println(counter)
//	lock.Unlock()
//}
//func main() {
//
//	lock := &sync.Mutex{}
//	for i := 0; i < 10; i++ {
//		go Count(lock)
//	}
//	for {
//		lock.Lock()
//		c := counter
//		lock.Unlock()
//		runtime.Gosched()
//		if c >= 10 {
//			break
//		}
//	}
//}
// 第3个例子
var counter int = 0

func Count(lock *sync.Mutex) {
	lock.Lock()

	counter++
	fmt.Println(counter)

	lock.Unlock()
}
func main() {
	// 并发执行引入了锁，每次改变counter，要等到解开才能执行下一个
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go Count(lock)
	}

	for {
		lock.Lock()

		c := counter
		fmt.Println(c, "x")

		lock.Unlock()

		runtime.Gosched()
		if c == 10 {
			break
		}
	}

}
