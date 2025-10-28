package main

import (
	"flag"
	"fmt"
	"strings"
)

type Namely struct {
	a int
}

type Namelys struct {
	b int
}

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "seperator")

func main() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}

	a := Namely{a: 10}
	b := Namelys{b: 20}

	if a.a == b.b{
		
	}
}
