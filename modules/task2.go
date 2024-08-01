package modules

import (
	"fmt"
	"sync"
)

func Fibonacci(ch chan<- int, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		close(ch)
	}()

	a, b := 0, 1
	for i := 0; i < 10; i++ {
		ch <- a
		a, b = b, a+b
	}
}

func GanjilGenap(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range ch {
		if num%2 == 0 {
			fmt.Println("Genap:", num)
		} else {
			fmt.Println("Ganjil:", num)
		}
	}
}