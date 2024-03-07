package main

import "fmt"

type Monkey struct {
	startingItems []int
	operation     func(int) int
	test          func(int) int
	inspect       int
	moduleFactor  int
}

func (m Monkey) String() string {
	items := len(m.startingItems)

	return fmt.Sprintf("Monkey has items: %d\t has Inspect: %d\n", items, m.inspect)
}
