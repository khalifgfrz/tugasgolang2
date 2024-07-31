package main

import "fmt"

func sum(d []int, ch chan int) {
	total := 0
	for _, v := range d {
        total += v
    }
    ch <- total
}

func main() {
    a := []int{7, 10, 2, 34, 33, -12, -8, 4}
    chn := make(chan int)
    
    go sum(a[:len(a)/2], chn)
    go sum(a[len(a)/2:], chn)
    
    result1 := <-chn
    result2 := <-chn
    
    fmt.Println("Sum of first half:", result1)
    fmt.Println("Sum of second half:", result2)
    fmt.Println("Total sum:", result1+result2)
}
