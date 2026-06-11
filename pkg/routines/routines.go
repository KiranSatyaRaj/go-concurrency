package routines

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// say prints each word of a phrase
func Say(wg *sync.WaitGroup, id int, phrase string) {
	for _, word := range strings.Fields(phrase) {
		fmt.Printf("Worker %#d says: %s...\n", id, word)
		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}
	wg.Done() // decreements WaitGroup counter by 1
}
