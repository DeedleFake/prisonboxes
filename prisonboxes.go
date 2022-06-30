package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func shuffle[T any](s []T) {
	rand.Shuffle(len(s), func(i1, i2 int) {
		s[i1], s[i2] = s[i2], s[i1]
	})
}

func check(boxes []int, prisoner int) bool {
	cur := prisoner
	for i := 0; i < (len(boxes)+1)/2; i++ {
		next := boxes[cur]
		if next == prisoner {
			return true
		}
		cur = next
	}
	return false
}

func simulate(boxes []int) bool {
	for p := range boxes {
		if !check(boxes, p) {
			return false
		}
	}
	return true
}

var boxPool sync.Pool

func getBoxes(np int) []int {
	boxes, _ := boxPool.Get().([]int)
	if boxes != nil {
		return boxes
	}

	boxes = make([]int, np)
	for i := range boxes {
		boxes[i] = i
	}
	return boxes
}

func putBoxes(boxes []int) {
	boxPool.Put(boxes)
}

func main() {
	np := flag.Int("p", 100, "number of prisoners")
	iter := flag.Int("n", 100000, "number of iterations to run")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	wg.Add(*iter)

	var success int64
	for i := 0; i < *iter; i++ {
		go func() {
			defer wg.Done()

			boxes := getBoxes(*np)
			defer putBoxes(boxes)

			shuffle(boxes)
			if simulate(boxes) {
				atomic.AddInt64(&success, 1)
			}
		}()
	}

	wg.Wait()

	fmt.Printf("Successful: %v\n", success)
	fmt.Printf("Success rate: %.2f\n", float64(success)*100/float64(*iter))
}
