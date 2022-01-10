package main

import "fmt"

func main() {
	colors := map[string]string{ //keys are string, values are string
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	// Other ways to create a map
	//var colors map[string]string
	//colors := make(map[string]string)
	// Deleting from a map
	//delete(colors, "white")
	colors["white"] = "#fffff"
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
