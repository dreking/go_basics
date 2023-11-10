package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	// var colors2 map[string]string

	// colors3 := make(map[string]string)

	// add new key value pair
	// colors2["white"] = "#ffffff"
	// colors3["black"] = "#000000"

	// delete(colors, "red")

	fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
