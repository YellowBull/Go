package main

import (
	"fmt"
	"runtime"
	"sync"
	//"time"
)

func main() {

	/**
	 * 基础goroutine
	 * @type {[type]}
	 */
	c := make(chan bool)
	go func() {
		fmt.Println("GO GO GO !")
		c <- true
		close(c) // 必须在正确的地方正确关闭，否则程序崩溃
	}()
	//<-c // 取出channel的值线程关闭
	for v := range c { // 循环遍历的方式取出channel的值
		fmt.Println(v) // true
	}

	/**
	 * runtime处理多核处理goroutine
	 */
	runtime.GOMAXPROCS(runtime.NumCPU())
	channel := make(chan bool, 10) // 指定要接受的channel个数
	for i := 0; i < 10; i++ {
		go add(channel, i)
	}
	for i := 0; i < 10; i++ {
		<-channel
	}

	fmt.Println("---------")

	/**
	 * sync处理多核goroutine
	 */
	wgroup := sync.WaitGroup{}
	wgroup.Add(10) // 添加处理事项个数
	for i := 0; i < 10; i++ {
		go countNum(&wgroup, i)
	}
	wgroup.Wait()

	/**
	 * select的使用
	 */
	ch := make(chan bool, 2)
	c1, c2 := make(chan int), make(chan string)
	go func() {
		for {
			select {
			case v, ok := <-c1: // 读出c1的值
				if !ok {
					ch <- true
					break
				}
				fmt.Println("c1:", v)
			case v, ok := <-c2: // 读出c2的值
				if !ok {
					ch <- true
					break
				}
				fmt.Println("c2:", v)
			}
		}
	}()

	c1 <- 1
	c2 <- "你好"
	c1 <- 2
	c2 <- "哈哈"

	close(c1)
	close(c2)

	for i := 0; i < 2; i++ {
		<-ch
	}

	/**
	 * select 产生 0/1 随机数
	 */
	//警告：小心尝试，这个随机数我是搞不出来，瞬间99%的CPU占用率
	random := make(chan int)
	go func() {
		for v := range random {
			fmt.Println(v)
		}
	}()

	for {
		select {
		case random <- 0:
		case random <- 1:
		case <-time.After(2 * time.Second):
			fmt.Println("TimeOut")
			break
		}
	}

}

func add(channel chan bool, index int) {
	a := 0
	for i := 1; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	channel <- true

}

func countNum(sy *sync.WaitGroup, index int) {
	a := 0
	for i := 1; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	sy.Done()
}
