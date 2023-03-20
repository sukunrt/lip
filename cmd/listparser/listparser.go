package main

import "github.com/sukunrt/lip/parser"

func main() {
	p := &parser.ListParser{}
	s := "[x, y, z] = [a, b, c]"
	p.Parse(s)
}
