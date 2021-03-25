package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {

	pool := NewPool()

	wg := new(sync.WaitGroup)

	wg.Add(1)
	// header render
	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(100 * time.Millisecond)

			render := &HeaderRender{
				Title: fmt.Sprintf("title-%d", i),
				Logo:  fmt.Sprintf("logo-%d", i),
			}

			if _, err := pool.Submit(render); err != nil {
				log.Printf("pool submit error: %v", err)
			}
		}

		wg.Done()
	}()

	// foot render
	wg.Add(1)
	go func() {
		for i := 0; i < 150; i++ {
			time.Sleep(80 * time.Millisecond)

			render := &FootRender{
				Statement: fmt.Sprintf("foot: %d", i),
				Year:      2000 + i,
			}

			_, err := pool.Submit(render)
			if err != nil {
				log.Printf("pool submit error: %v", err)
			}
		}

		wg.Done()
	}()

	wg.Wait()
}
