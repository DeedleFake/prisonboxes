package main

import (
	"flag"
	"fmt"
	"math/rand"
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

func main() {
	np := flag.Int("p", 100, "number of prisoners")
	iter := flag.Int("n", 100000, "number of iterations to run")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	boxes := make([]int, *np)
	for i := range boxes {
		boxes[i] = i
	}

	var success int
	for i := 0; i < *iter; i++ {
		shuffle(boxes)
		if simulate(boxes) {
			success++
		}
	}

	fmt.Printf("Successful: %v\n", success)
	fmt.Printf("Success rate: %.2f\n", float64(success)*100/float64(*iter))
}
