package main

import (
	"fmt"
	"khalifgfrz/tugasgolang2/modules"
)

func main() {
    a := []int{7, 10, 2, 34, 33, -12, -8, 4}
    chn := make(chan int)
    
    go modules.Sum(a[:len(a)/2], chn)
    go modules.Sum(a[len(a)/2:], chn)
    
    result1 := <-chn
    result2 := <-chn
    
    fmt.Println("Sum of first half:", result1)
    fmt.Println("Sum of second half:", result2)
    fmt.Println("Total sum:", result1+result2)

	// var wg sync.WaitGroup
	// ch := make(chan int)

	// wg.Add(2)
	// go modules.Fibonacci(ch, &wg)
	// go modules.GanjilGenap(ch, &wg)

	// wg.Wait()

	// distanceChan := make(chan int)

	// wg.Add(2)
	// go modules.TotalPrice(distanceChan, &wg)
	// go modules.SendDistance(distanceChan, 10, &wg)

	// wg.Wait()
}
