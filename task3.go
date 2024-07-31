package main

import (
	"fmt"
	"sync"
)

func calculatePrice(distance int) int {
	const pricePerKm = 2000
	return distance * pricePerKm
}

func totalPrice(distanceChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	totalPrice := 0
	for distance := range distanceChan {
		price := calculatePrice(distance)
		totalPrice += price
	}
	fmt.Printf("Total harga untuk jarak tempuh yang diberikan: %d\n", totalPrice)
}

func sendDistance(distanceChan chan<- int, distance int, wg *sync.WaitGroup) {
	defer wg.Done()
	distanceChan <- distance
	close(distanceChan)
}

func main() {
	var wg sync.WaitGroup

	distanceChan := make(chan int)

	wg.Add(2)
	go totalPrice(distanceChan, &wg)
	go sendDistance(distanceChan, 10, &wg)

	wg.Wait()
}
