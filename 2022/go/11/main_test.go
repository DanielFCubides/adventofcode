package main

import (
	"testing"
)

func TestMonkeyInTheMiddle(t *testing.T) {
	t.Run("Validate example", func(t *testing.T) {
		monkeys := MonkeyInTheMiddle(createMonkeysExample())
		if monkeys[0].startingItems[0] != 20 {
			t.Errorf("Output expected is :\n\nMonkey 0: 20, 23, 27, 26\nMonkey 1: 2080, 25, "+
				"167, 207, 401, 1046\nMonkey 2: \nMonkey 3: \n\n resulted : \n %v", monkeys)
		}
	})

}
