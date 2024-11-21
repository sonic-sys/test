package main

import (
	"fmt"
)

func main() {
	arr := []int{22, 33, 44, 1, 1, 3, 4, 5}
	result := make(map[int]int)

	for _, a := range arr {
		result[a]++
	}

	for key, value := range result {
		fmt.Printf("число %d == %d раз(а)\n", key, value)
	}
}
