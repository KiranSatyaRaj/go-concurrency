package digitcount

import (
	"strings"
	"sync"
	"unicode"
)

func countDigits(str string) int {
	count := 0

	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}

	return count
}

// this is an exercise
// link: https://antonz.org/go-concurrency/goroutines/
// exercise name: counting digits in words
func CountDigitsInWords(phrase string) int {
	var wg sync.WaitGroup
	syncStats := new(sync.Map)
	words := strings.Fields(phrase)

	for _, word := range words {
		var count int
		wg.Go(func() {
			count = countDigits(word)
			syncStats.Store(word, count)
		})
		wg.Wait()
	}
	return 0
}
