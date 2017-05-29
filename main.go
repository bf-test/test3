package main

import "fmt"

func main() {
	fmt.Println("vim-go %s", "hello") // govet: possible formatting directive in Println call
}
