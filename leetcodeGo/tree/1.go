package main

import (
	"fmt"
	"sync"
)

// 模拟订单对象
type OrderInfo struct {
	id int
}

var num int
var wg sync.WaitGroup

// 生产订单--生产者
func producer(out chan<- OrderInfo) {
	// 业务生成订单
	for i := 0; i < 10000; i++ {
		order := OrderInfo{id: i + 1}
		fmt.Println("生成订单，订单ID为：", order.id)
		out <- order // 写入channel
	}
	// 如果不关闭，消费者就会一直阻塞，等待读
	close(out) // 订单生成完毕，关闭channel
}

func producerDispatch(tasks chan OrderInfo) {

	producer(tasks)
}

// 处理订单--消费者
func consumer(i int, in <-chan OrderInfo) {
	defer wg.Done()
	// 从channel读取订单，并处理
	for order := range in {
		fmt.Println("消费者", i, "读取订单,订单ID为:", order.id)
		num++
	}
}

func main() {
	//消费者个数
	var channelLen = 50

	tasks := make(chan OrderInfo, channelLen)

	go producerDispatch(tasks)
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		go consumer(i, tasks)
	}

	wg.Wait()
	//defer close(tasks)
	fmt.Println("all done")
	fmt.Println(num)
}
