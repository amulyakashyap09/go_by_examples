package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	logCh := make(chan string, 50)
	mutex := make(chan bool, 1)
	defer close(logCh)
	defer close(mutex)
  
	file := "./log.txt"
	f, _ := os.Create(file)

	go func() {
		for {
			msg, ok := <-logCh
			if ok {
				f, _ = os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)
				logTime := time.Now().Format(time.RFC3339)
				f.WriteString(logTime + " - " + msg)
				f.Close()
				wg.Done()
			} else {
				break
			}
		}
	}()
	for i := 0; i <= 10; i++ {
		for j := 0; j <= 10; j++ {
			wg.Add(1)
			mutex <- true
			go func(x int, y int) {
				msg := fmt.Sprintf("%d + %d = %d\n", x, y, x+y)
				logCh <- msg
				<-mutex
			}(i, j)
		}
	}

	// time.Sleep(10 * time.Millisecond)
	fmt.Println("done!!!")
	wg.Wait()
}
