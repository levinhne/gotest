package main

import (
	"fmt"
	"log"
	"runtime/debug"
)

type Parent struct {
	Name  string
	Child *struct {
		Name string
	}
}

func main() {
	defer func() {
		if x := recover(); x != nil {
			println(fmt.Sprintf("%T: %+v", x, x))
			debug.PrintStack()
		}
	}()
	var p *Parent
	p.Child = nil
	log.Println(p.Child.Name)
}
