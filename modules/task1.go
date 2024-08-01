package modules

func Sum(d []int, ch chan int) {
	total := 0
	for _, v := range d {
        total += v
    }
    ch <- total
}