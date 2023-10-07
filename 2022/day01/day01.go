package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func mostCalories(calories []int) int {
	sort.Ints(calories)
	sum := 0
	for i, count := len(calories)-1, 0; count < 3; i, count = i-1, count+1 {
		sum += calories[i]
	}
	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var calories []int
	sum := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "stop" {
			break
		}
		num, err := strconv.Atoi(text)
		if err != nil {
			calories = append(calories, sum)
			sum = 0
			continue
		}
		sum += num
	}
	fmt.Println(mostCalories(calories))
}
