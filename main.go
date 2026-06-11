package main

import (
	"fmt"
	"sync"

	"github.com/KiranSatyaRaj/go-concurrency/pkg/routines"
)

func main() {
	var wg sync.WaitGroup // has a counter variable

	wg.Add(2) // increments the WaitGroup counter variable by two
	// do not mix business logic with concurrent logic
	go func() {
		defer wg.Done()
		routines.Say(&wg, 1, "go is awesome")
	}()

	go func() {
		defer wg.Done()
		routines.Say(&wg, 2, "guns of ganymede")
	}()

	wg.Wait() // waits until counter is zero here main has to wait

	// ALternatively we can use waitGroup.Go()

	wg.Go(func() {
		fmt.Println("go is awesome")
	})

	wg.Go(func() {
		fmt.Println("guns of ganymede")
	})

	wg.Wait()
	fmt.Println("Done")
}
