package main

import "fmt"

func main() {
	for i :=  50000; i < 50100; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}
}
