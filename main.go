package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// letterWorker represents a worker that specializes in delivering one letter
func letterWorker(letter rune, position int, letterChan chan<- struct {
	char     rune
	position int
}, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate "processing time" for each letter
	time.Sleep(time.Duration(position) * 100 * time.Millisecond)

	// Send the letter and its position through the channel
	letterChan <- struct {
		char     rune
		position int
	}{letter, position}
}

func main() {
	message := "Hello, World!"
	letters := []rune(message)

	// Create a channel to receive letters
	letterChan := make(chan struct {
		char     rune
		position int
	}, len(letters))

	// Create a WaitGroup to synchronize our letter workers
	var wg sync.WaitGroup

	// Launch a goroutine for each letter
	for i, letter := range letters {
		wg.Add(1)
		go letterWorker(letter, i, letterChan, &wg)
	}

	// Launch a goroutine to close the channel after all workers are done
	go func() {
		wg.Wait()
		close(letterChan)
	}()

	// Create a slice to store our final message
	result := make([]string, len(letters))

	// Collect letters as they arrive and build our message
	for letter := range letterChan {
		result[letter.position] = string(letter.char)

		// Print the current state of the message
		fmt.Printf("\r%s%s",
			strings.Join(result, ""),
			strings.Repeat(" ", len(letters)-letter.position-1))
	}
}
