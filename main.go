package main

import (
	"fmt"
	"util/rand"
)

func main() {
	a := rand.GenerateUniqueID(1, 10)
	fmt.Println(a)
}
