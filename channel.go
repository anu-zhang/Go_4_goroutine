package main

import (
	"fmt"
)

func ChannelCount(ch chan int) {
	//在对应的channel写入 数据前，这个操作也是阻塞的。
	ch <- 1 // 这里阻塞 ， 读完之后 才会进行下一步打印操作
	fmt.Println("Counting")
}

func OtherExsample()  {
	ch := make(chan int, 1)
	for {
		select {
		case ch <- 10:
		case ch <- 1:
		}
		i := <-ch
		fmt.Println("Value received:", i)
	}
}
func main() {
	// 获得cpu核心数
	//num := runtime.NumCPU()
	//fmt.Println(num)

	fmt.Println(1);
	chs := make([]chan int, 10)
	fmt.Println(2)
	for i := 0; i < 10; i++ {
		fmt.Println("for")
		chs[i] = make(chan int)
		fmt.Println("for2")
		go ChannelCount(chs[i]) //传递的应该是地址
		fmt.Println("for3")
	}

	fmt.Println("3")
	for _, ch := range (chs) {
		fmt.Println("range")
		<-ch // 只有当最后一个读取完，才会打印 ChannelCount() 方法的 Counting
		fmt.Println("range2")
	}
	fmt.Println(4)
}
