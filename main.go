package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// asciiLetters contains the ASCII art representation for each character
var asciiLetters = map[rune][]string{
	'H': {
		"█   █",
		"█   █",
		"█████",
		"█   █",
		"█   █",
	},
	'e': {
		"█████",
		"█    ",
		"█████",
		"█    ",
		"█████",
	},
	'l': {
		"█    ",
		"█    ",
		"█    ",
		"█    ",
		"█████",
	},
	'o': {
		"█████",
		"█   █",
		"█   █",
		"█   █",
		"█████",
	},
	',': {
		"     ",
		"     ",
		"     ",
		"  █  ",
		" █   ",
	},
	' ': {
		"     ",
		"     ",
		"     ",
		"     ",
		"     ",
	},
	'W': {
		"█   █",
		"█   █",
		"█ █ █",
		"██ ██",
		"█   █",
	},
	'r': {
		"█████",
		"█   █",
		"█████",
		"█  █ ",
		"█   █",
	},
	'd': {
		"█████",
		"█   █",
		"█   █",
		"█   █",
		"█████",
	},
	'!': {
		"  █  ",
		"  █  ",
		"  █  ",
		"     ",
		"  █  ",
	},
}

// letterWorker now delivers ASCII art letters instead of single characters
func letterWorker(letter rune, position int, letterChan chan<- struct {
	art      []string
	position int
}, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate "processing time" for each letter
	time.Sleep(time.Duration(position) * 100 * time.Millisecond)

	// Send the ASCII art and position through the channel
	letterChan <- struct {
		art      []string
		position int
	}{asciiLetters[letter], position}
}

func main() {
	message := "Hello, World!"
	letters := []rune(message)

	// Create a channel to receive ASCII art letters
	letterChan := make(chan struct {
		art      []string
		position int
	}, len(letters))

	var wg sync.WaitGroup

	// Launch a goroutine for each letter
	for i, letter := range letters {
		wg.Add(1)
		go letterWorker(letter, i, letterChan, &wg)
	}

	// Close channel after all workers finish
	go func() {
		wg.Wait()
		close(letterChan)
	}()

	// Create a matrix to store our ASCII art message
	// Each character is 5 lines tall and 6 characters wide (5 + 1 space between letters)
	result := make([][]string, 5) // 5 rows
	for i := range result {
		result[i] = make([]string, len(letters))
		for j := range result[i] {
			result[i][j] = strings.Repeat(" ", 6)
		}
	}

	// Collect and display ASCII letters as they arrive
	letterCount := 0
	for letter := range letterChan {
		// Add the ASCII art to our result matrix
		for i, line := range letter.art {
			result[i][letter.position] = line + " "
		}

		// Clear screen (basic approach) and display current state
		fmt.Print("\033[H\033[2J")
		for _, line := range result {
			fmt.Println(strings.Join(line, ""))
		}

		letterCount++
	}
}
