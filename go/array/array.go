package main

import "fmt"

var numbers = [100]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func read(index int) int {
	return numbers[index]
}
func search(search int) int {
	var result int
	for index, number := range numbers {
		if search == number {
			result = index
		}
	}
	return result
}
func insertBeginning() {}
func insertend()       {}
func deleteBeginning() {}
func deleteend()       {}

func main() {
	fmt.Println(read(0))
	fmt.Println(search(4))
}
