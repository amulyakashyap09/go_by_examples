package main

import (
	"fmt"
)

type PurchaseOrder struct {
	Number int
	Value  float64
}

func savePo(po *PurchaseOrder, callback chan *PurchaseOrder) {
	po.Number = 100
	po.Value = 120
	callback <- po
}

func f(n int, callback chan int) {

	var sum int = 0
	for i := 0; i < 10; i++ {
		fmt.Println("hi ...", i)
		sum += 1
	}
	callback <- sum
}

func main() {
	po := new(PurchaseOrder)
	cb := make(chan *PurchaseOrder)
	go savePo(po, cb)
	
	sumCh := make(chan int)
	defer close(sumCh)

	go f(0, sumCh)

	data := <-cb
	fmt.Println(data, <-sumCh)
}
