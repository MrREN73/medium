package main

import "fmt"

func main() {

	data1 := []int{1, 24, 56, 22, 57, 3, 74, 34, 26, 99, 32, 53, 68, 90, 34, 28, 75, 21}
	data2 := []int{34, 23, 67, 22, 86, 34, 97, 93, 23, 27, 43, 23, 45, 78, 3, 24, 12, 36}

	// TODO calc medium value (43.52) using fan-in fan-out

	ch1 := make(chan int)
	ch2 := make(chan int)
	out := make(chan float32)
	done := make(chan bool)

	go calc(ch1, ch2, out, done)

	for _, v := range data1 {
		ch1 <- v
	}

	for _, v := range data2 {
		ch2 <- v
	}

	done <- true

	fmt.Println(<-out)
}

func calc(in1, in2 chan int, out chan float32, done chan bool) {
	sum := 0
	count := 0
	for {
		select {
		case v1 := <-in1:
			sum += v1
			count++
		case v2 := <-in2:
			sum += v2
			count++
		case <-done:
			out <- float32(sum) / float32(count)
			return
		}
	}
}
