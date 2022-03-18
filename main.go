package main

import (
	"fmt"
)

func main() {
	greet("Berkay")
}

func greet(name string) {
	fmt.Printf("Selam %s :)\n", name)
}
