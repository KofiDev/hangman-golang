package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	playing := true
	symbol := "_"

	vehicles := []string{"car", "boat", "bike", "ship", "plane"}
	animals := []string{"cat", "dog", "elephant", "horse", "hawk"}
	foods := []string{"spaghetti", "apple", "sandwich", "cheesecake", "salad"}

	categories := [][]string{vehicles[:], animals[:], foods[:]}

	for playing {
		rand.Seed(time.Now().UnixNano())
		randomCategoryIndex := rand.Intn(len(categories))
		randomCategory := categories[randomCategoryIndex]

		randomWord := randomCategory[rand.Intn(len(randomCategory))]

		switch randomCategoryIndex {
		case 0:
			fmt.Println("Category: vehicles")
		case 1:
			fmt.Println("Category: animals")
		case 2:
			fmt.Println("Category: foods")
		}
		fmt.Printf("The word is: %s\n", randomWord)

		fmt.Println(strings.Repeat(symbol, len(randomWord)))
		playing = false
	}
}
