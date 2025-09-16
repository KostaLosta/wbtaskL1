package main

import (
	"fmt"
	"sync"
)

func main() {

	array := [5]int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	result := make([]int, len(array))

	for i, v := range array {
		wg.Add(1)
		go func(index, value int) {
			defer wg.Done()
			result[index] = value * value
		}(i, v)
	}

	wg.Wait()

	for _, r := range result {
		fmt.Println(r)
	}
}
