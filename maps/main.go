package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"red":   "ABC",
	// 	"green": "XYZ",
	// }

	// var colors map[string]string
	colors := make(map[string]string)
	// fmt.Println(colors)
	colors["red"] = "PQR"
	colors["green"] = "ABC"

	// fmt.Println(colors)

	// delete(colors, "red")
	// fmt.Println(colors)
	printMap(colors)

}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println(key + " " + value)
	}
}
