package main

import "fmt"

func main() {
	fmt.Println(Hello("Hello"))
}

func Hello(name string) string {
	return "Hello " + name
}
