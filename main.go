package main

import (
	"daniel/golang-parallel-sort/utils"
	"fmt"
)

const size = 1000000

func main() {
	s := utils.Random(size)
	fmt.Println(s)
}
