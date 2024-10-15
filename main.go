package main

import (
	"fmt"
	"github.com/yungchoyee/cutil/rand"
)

func main() {
	a := rand.GenerateUniqueID(1, 10)
	fmt.Println(a)
}
