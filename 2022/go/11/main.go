package main

import "fmt"

func InspectItemPart1(m *Monkey, monkeys []*Monkey, divisibleFactor int) {
	item := m.startingItems[0]
	m.startingItems = m.startingItems[1:]
	//newWorryLevel := m.operation(item % m.moduleFactor)
	newWorryLevel := m.operation(item)
	newWorryLevel = (int)(newWorryLevel / 3)
	throwTo := m.test(newWorryLevel)
	monkeys[throwTo].startingItems = append(monkeys[throwTo].startingItems, newWorryLevel)
	m.inspect++
}

func InspectItemPart2(m *Monkey, monkeys []*Monkey, divisibleFactor int) {
	item := m.startingItems[0]
	m.startingItems = m.startingItems[1:]
	newWorryLevel := m.operation(item % divisibleFactor)
	throwTo := m.test(newWorryLevel)
	monkeys[throwTo].startingItems = append(monkeys[throwTo].startingItems, newWorryLevel)
	m.inspect++
}
func InspectItems(monkeyid int, monkeys []*Monkey, InspectItem func(m *Monkey, monkeys []*Monkey, divisibleFactor int)) {
	m := monkeys[monkeyid]
	numberOfItems := len(m.startingItems)
	divisibleFactor := calculateDivisibleProduct(monkeys)
	if numberOfItems > 0 {
		for i := 0; i < numberOfItems; i++ {
			InspectItem(
				m,
				monkeys,
				divisibleFactor,
			)
		}
	}

}

func calculateDivisibleProduct(monkeys []*Monkey) int {
	divisibleProduct := 1
	for _, monkey := range monkeys {
		divisibleProduct = divisibleProduct * monkey.moduleFactor
	}
	return divisibleProduct
}

func MonkeyBusiness(monkeys []*Monkey) int {
	MostActiveMonkey, secondMostActiveMonkey := monkeys[0].inspect, monkeys[1].inspect
	for i := 2; i < len(monkeys); i++ {
		monkeyActivity := monkeys[i].inspect
		if monkeyActivity > MostActiveMonkey && monkeyActivity > secondMostActiveMonkey {
			secondMostActiveMonkey = MostActiveMonkey
			MostActiveMonkey = monkeyActivity
			continue
		}
		if monkeyActivity < MostActiveMonkey && monkeyActivity > secondMostActiveMonkey {
			secondMostActiveMonkey = monkeyActivity

		}
	}
	return MostActiveMonkey * secondMostActiveMonkey
}

func MonkeyInTheMiddle(monkeys []*Monkey, rounds int, InspectItem func(m *Monkey, monkeys []*Monkey, divisibleFactor int)) ([]*Monkey, int) {
	for i := 0; i < rounds; i++ {
		//fmt.Println("Round ", i+1)
		for monkeyid, _ := range monkeys {
			InspectItems(monkeyid, monkeys, InspectItem)
		}
	}
	fmt.Println(monkeys)
	monkeyBusiness := MonkeyBusiness(monkeys)
	return monkeys, monkeyBusiness
}

func main() {
	monkeys := createMonkeysInput()
	_, monkeyBusiness := MonkeyInTheMiddle(monkeys, 10000, InspectItemPart2)
	fmt.Println(monkeyBusiness)

}
