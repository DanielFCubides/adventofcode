package main

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
		inspect:      0,
		moduleFactor: 23,
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
			inspect:      0,
			moduleFactor: 19,
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
			inspect:      0,
			moduleFactor: 13,
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
			inspect:      0,
			moduleFactor: 17,
		}}
}

func createMonkeysInput() []*Monkey {
	return []*Monkey{
		{
			startingItems: []int{76, 88, 96, 97, 58, 61, 67},
			operation: func(i int) int {
				return i * 19
			},
			test: func(i int) int {
				if i%3 == 0 {
					return 2 // throw to monkey 2
				}
				return 3 // throw to monkey 3
			},
			inspect:      0,
			moduleFactor: 3,
		},
		{
			startingItems: []int{93, 71, 79, 83, 69, 70, 94, 98},
			operation: func(i int) int {
				return i + 8
			},
			test: func(i int) int {
				if i%11 == 0 {
					return 5 // throw to monkey 5
				}
				return 6 // throw to monkey 6
			},
			inspect:      0,
			moduleFactor: 11,
		},
		{
			startingItems: []int{50, 74, 67, 92, 61, 76},
			operation: func(i int) int {
				return i * 13
			},
			test: func(i int) int {
				if i%19 == 0 {
					return 3 // throw to monkey 3 if divisible by 19
				}
				return 1 // throw to monkey 1 if not divisible by 19
			},
			inspect:      0,
			moduleFactor: 19,
		},
		{
			startingItems: []int{76, 92},
			operation: func(i int) int {
				return i + 6
			},
			test: func(i int) int {
				if i%5 == 0 {
					return 1 // throw to monkey 1 if divisible by 5
				}
				return 6 // throw to monkey 6 if not divisible by 5
			},
			inspect:      0,
			moduleFactor: 5,
		},
		{
			startingItems: []int{74, 94, 55, 87, 62},
			operation: func(i int) int {
				return i + 5
			},
			test: func(i int) int {
				if i%2 == 0 {
					return 2 // throw to monkey 2 if divisible by 2
				}
				return 0 // throw to monkey 0 if not divisible by 2
			},
			inspect:      0,
			moduleFactor: 2,
		},
		{
			startingItems: []int{59, 62, 53, 62},
			operation: func(i int) int {
				return i * i // Squaring the item
			},
			test: func(i int) int {
				if i%7 == 0 {
					return 4 // throw to monkey 4 if divisible by 7
				}
				return 7 // throw to monkey 7 if not divisible by 7
			},
			inspect:      0,
			moduleFactor: 7,
		},
		{
			startingItems: []int{62},
			operation: func(i int) int {
				return i + 2
			},
			test: func(i int) int {
				if i%17 == 0 {
					return 5 // throw to monkey 5 if divisible by 17
				}
				return 7 // throw to monkey 7 if not divisible by 17
			},
			inspect:      0,
			moduleFactor: 17,
		},
		{
			startingItems: []int{85, 54, 53},
			operation: func(i int) int {
				return i + 3
			},
			test: func(i int) int {
				if i%13 == 0 {
					return 4 // throw to monkey 4 if divisible by 13
				}
				return 0 // throw to monkey 0 if not divisible by 13
			},
			inspect:      0,
			moduleFactor: 13,
		},
	}
}
