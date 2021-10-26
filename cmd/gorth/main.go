package main

import (
	"flag"
	"fmt"

	"github.com/kshmatov/gorth/internal/runner"
)

func main() {
	fn := flag.String("f", "", "file with sourcecode")
	debug := flag.Bool("d", false, "show debug output")

	flag.Parse()

	if *fn == "" {
		flag.PrintDefaults()
		return
	}
	err := runner.Run(*fn, *debug)
	if err != nil {
		fmt.Printf("error:\n\t%v\n", err)
	} else {
		fmt.Printf("\ndone ok\n")
	}
}
