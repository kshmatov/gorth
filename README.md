# gorth

Forth'like language with Go, just to try

I don't care about optimization at all, this is my first attempt to make interpreter/compiler

## Howto

#### Use Makefile, Luke!

To get dependencies

```shell
make deps
```

Tests

```shell
make test
```

Build

```shell
make build
```

Run sample/test.gorth in debug mode

```shell
make run
```

Default __make__ will use deps, test, build

## Usage

```shell
gorth [-d] -f filename
```

use **-d** for debug output

## Language syntax

Only integers (64 bit)

Number will push it to stack

**add**, **sub**, **+**, **-** will pop two args from stack, does it and push the result to stack

**dump**, **.** will pop one arg from stack and print it

Space, tab or new line: separate operands

**45 34 +** is correct sample

**45 34+** will rise syntax error

I plan to add arithmetics, bitwise operations, strings, variables, conditions and cicles. May be.

Also I plan to make virtual machine for it and compiler for this VM

## Samples

1. test.gorth - normal program
2. syntax.gorth - syntax error in code
3. outstack.gorth - runtime error
4. maxint.gorth - sum of two max int64

Work in progress