package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("Doing task", id)
	// w.Done()
}

func main() {
	var wg sync.WaitGroup
	for i := range 10 {
		wg.Add(1)
		go task(i, &wg)

		// go func(id int) {
		// 	task(id)
		// }(i)

	}
	wg.Wait()
}
