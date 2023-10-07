package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Rock - 1 Papper - 2 Scissors - 3
	scanner := bufio.NewScanner(os.Stdin)
	points := rune(0)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "stop" {
			break
		}
		symbols := []rune(text)
		enemy := symbols[0] - 'A' + 1
		player := symbols[2] - 'X' + 1
		points += player
		if enemy == player {
			points += 3
			continue
		}
		if player == 1 && enemy == 3 ||
			player == 2 && enemy == 1 ||
			player == 3 && enemy == 2 {
			points += 6
		}
	}
	fmt.Println(points)
}
