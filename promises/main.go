package main

import (
	"errors"
	"fmt"
	"time"
)

type PurchaseOrder struct {
	Number int
	Value  float64
}

func savePo(po *PurchaseOrder, itShouldFail bool) *Promise {
	result := new(Promise)
	result.successChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error, 1)

	go func() {
		if itShouldFail {
			result.successChannel <- errors.New("Failed to save orders")
		} else {
			po.Number = 13254
			po.Value = 123562537
			result.successChannel <- po
		}
	}()
	return result
}

func main() {
	po := new(PurchaseOrder)
	savePo(po, false).Then(func(obj interface{}) error {
		po := obj.(*PurchaseOrder)
		fmt.Println("Purchase Order : ", po.Number, po.Value)
		return nil
	}, func(err error) {
		fmt.Println(err)
	})

	// fmt.Scanln()
	time.Sleep(10 * time.Second)
}

type Promise struct {
	successChannel chan interface{}
	failureChannel chan error
}

func (this *Promise) Then(success func(interface{}) error, failure func(error)) *Promise {
	result := new(Promise)
	result.successChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error, 1)

	go func() {
		select {
		case obj := <-this.successChannel:
			newError := success(obj)
			if newError == nil {
				result.successChannel <- obj
			} else {
				result.failureChannel <- newError
			}
		case err := <-this.failureChannel:
			failure(err)
			result.failureChannel <- err
		}
	}()

	return result
}
