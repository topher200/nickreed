package main

import (
	"bufio"
	"fmt"
	"os"
)

// promptForCommand shows the user a menu and returns their response.
//
// Returns a rune with the user's input.
func promptForCommand(consoleReader *bufio.Reader) (rune, error) {
	fmt.Println("Options:")
	fmt.Println(" 1: Enter a '1'")
	fmt.Println(" 2: Enter a '-1'")
	fmt.Println(" r: Reset")
	fmt.Println(" q: Quit")

	// We loop our reader so we can ignore newlines
	for {
		rune, _, err := consoleReader.ReadRune()
		if rune != '\n' {
			return rune, err
		}
	}
}

type Score struct {
	tally, totalClicks int
}

func runLoop() {
	consoleReader := bufio.NewReader(os.Stdin)
	score := Score{}
	for {
		fmt.Printf("Tally: %2d, Total: %2d\n", score.tally, score.totalClicks)
		command, _ := promptForCommand(consoleReader)
		switch command {
		case '1':
			score.tally += 1
			score.totalClicks += 1
		case '2':
			score.tally += -1
			score.totalClicks += 1
		case 'r':
			fmt.Println("Reset!")
			score = Score{}
		case 'q':
			return
		default:
			fmt.Printf("Unknown command: %q\n", command)
		}
	}
}

func main() {
	runLoop()
	fmt.Println("Done!")
}
