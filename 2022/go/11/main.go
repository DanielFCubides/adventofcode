package main

import "fmt"

type Monkey struct {
	startingItems []int
	operation     func(int) int
	test          func(int) int
}

func (m Monkey) String() string {
	items := m.startingItems
	return fmt.Sprintf("\tmonkey starting items: %#v\n", items)
}

func InspectItems(monkeyid int, monkeys []*Monkey) {
	m := monkeys[monkeyid]
	numberOfItems := len(m.startingItems)
	fmt.Printf("monkey %d has %d items\n", monkeyid, numberOfItems)
	if numberOfItems > 0 {
		for i := 0; i < numberOfItems; i++ {
			InspectItem(m, monkeys)
		}
	}

}

func InspectItem(m *Monkey, monkeys []*Monkey) {
	item := m.startingItems[0]
	fmt.Printf("\tMonkey inspects an item with a worry level of %d\n", item)
	m.startingItems = m.startingItems[1:]
	newWorryLevel := m.operation(item)
	newWorryLevel = (int)(newWorryLevel / 3)
	fmt.Printf("\tMonkey gets bored with item. Worry level is divided by 3 to %d\n", newWorryLevel)
	throwTo := m.test(newWorryLevel)
	fmt.Printf("\tItem with worry level %d is thrown to monkey %d\n\n", newWorryLevel, throwTo)
	monkeys[throwTo].startingItems = append(monkeys[throwTo].startingItems, newWorryLevel)
}

func main() {
	monkeys := createMonkeysExample()
	fmt.Println(monkeys)
	MonkeyInTheMiddle(monkeys)

}

func MonkeyInTheMiddle(monkeys []*Monkey) []*Monkey {
	for i := 0; i < 1; i++ {
		for monkeyid, _ := range monkeys {
			fmt.Printf("Monkey %d:\n", monkeyid)
			InspectItems(monkeyid, monkeys)
		}
		fmt.Println(monkeys)
	}
	return monkeys
}

func createMonkeysExample() []*Monkey {
	return []*Monkey{{
		startingItems: []int{79, 98},
		operation: func(i int) int {
			return i * 19
		},
		test: func(i int) int {
			if i%23 == 0 {
				return 2
			}
			return 3
		},
	},
		{
			startingItems: []int{54, 65, 75, 74},
			operation: func(i int) int {
				return i + 6
			},
			test: func(i int) int {
				if i%19 == 0 {
					return 2
				}
				return 0
			},
		},
		{
			startingItems: []int{79, 60, 97},
			operation: func(i int) int {
				return i * i
			},
			test: func(i int) int {
				if i%13 == 0 {
					return 1
				}
				return 3
			},
		},
		{
			startingItems: []int{74},
			operation: func(i int) int {
				return i + 3
			},
			test: func(i int) int {
				if i%17 == 0 {
					return 0
				}
				return 1
			},
		}}
}
