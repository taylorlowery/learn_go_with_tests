package main

import "fmt"

const englishHelloPrefix = "Hello, "

func greet(name string) string {
	return "Howdy, " + name + "!"
}

func main() {
	fmt.Println(greet("Tater"))
}
