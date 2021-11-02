package interpreter

import (
	"fmt"
	"os"

	"github.com/kshmatov/gorth/internal/programm"
)

func Simulate(prog programm.Programm, debug bool) error {
	defer func() {
		if e := recover(); e != nil {
			//err = errors.Wrapf(e.(error), "line [%v], %v", i, operation)
		}
	}()
	out := make(chan string)
	go func() {
		for v := range out {
			if debug {
				fmt.Fprintln(os.Stdout, v)
			}
		}
	}()

	return prog.Run(out)
}

func debugOut(s string) {
	_, err := fmt.Fprint(os.Stderr, s)
	if err != nil {
		fmt.Print(s)
	}
}
