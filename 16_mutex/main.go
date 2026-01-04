package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) incrementViews(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()

	p.mu.Lock()
	p.views++
}

func main() {
	myPost := post{views: 0}
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		wg.Go(func() {
			myPost.incrementViews(&wg)
		})
	}

	wg.Wait()

	fmt.Println(myPost.views)
}
