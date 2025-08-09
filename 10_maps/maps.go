package main

import "fmt"

func main() {
	// creating map

	m := make(map[string]string)
	m["name"] = "Ritochit Ghosh"
	m["role"] = "Backend developer"

	fmt.Println(m["name"], m["role"])
}
