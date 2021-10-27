package main

import (
	"flag"
	"fmt"
	"os"

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
		fmt.Fprintln(os.Stderr, err.Error())
	} else {
		fmt.Println("ok")
	}
}
