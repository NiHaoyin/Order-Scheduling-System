package main

import (
	"fmt"
)

func main() {
	mapping := make(map[int][]int)
	slice := make([]int, 10)
	mapping[0] = slice
	slice[0] = 9
	//slice = append(slice, 100)
	//mapping[0] = slice
	fmt.Println(mapping)

}
