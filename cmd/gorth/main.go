package main

import (
	"fmt"

	"github.com/kshmatov/gorth/internal/op"
)

func main() {
	fmt.Println("Starting gorth")
	fmt.Printf("%#v %T\n", op.COUNT_OPS, op.COUNT_OPS)
}
