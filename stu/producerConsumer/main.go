package main

import (
	"fmt"
)

type Order struct {
	id int
}

func producer(out chan<- Order, num int) {
	for i := 0; i < num; i++ {
		fmt.Println("产生订单号：", i+1)
		out <- Order{id: i + 1}

	}
	close(out)
}
func consumer(in <-chan Order) {
	for order := range in {
		fmt.Println("用户处理订单：", order.id)
	}
}
func consumer2(in <-chan Order) {
	for {
		order, ok := <-in
		if !ok {
			return
		}
		fmt.Println("2处理订单：", order.id)
	}
}
func main() {
	order := make(chan Order, 5)
	go producer(order, 10)
	//consumer(order)
	consumer2(order)

}
