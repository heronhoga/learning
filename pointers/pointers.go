package main

import "fmt"

func main() {
	text1 := "hello"
	text2 := &text1
	text3 := *text2
	fmt.Println("text 1 val : ", text1)
	fmt.Println("text 2 val : ", text2)
	fmt.Println("text 3 val : ", text3)
}