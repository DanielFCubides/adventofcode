package main

import (
	"fmt"
	"testing"
)

func TestMonkeyInTheMiddlePart1(t *testing.T) {
	var tests = []struct {
		name     string
		expected []int
		round    int
		result   int
	}{
		{
			name:     "Round 1",
			expected: []int{2, 4, 3, 5},
			round:    1,
			result:   20,
		}, {
			name:     "Round 20",
			expected: []int{101, 95, 7, 105},
			round:    20,
			result:   10605,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			monkeys, _ := MonkeyInTheMiddle(createMonkeysExample(), tt.round, InspectItemPart1)
			expected := tt.expected
			output := []int{monkeys[0].inspect, monkeys[1].inspect, monkeys[2].inspect, monkeys[3].inspect}
			outputResult := fmt.Sprintf("%v", output)
			expectedResult := fmt.Sprintf("%v", expected)
			if outputResult != expectedResult {
				t.Fatalf("\nexpecetd %s \n got %s", expectedResult, outputResult)
			}

		})
	}

}

func TestMonkeyInTheMiddlePart2(t *testing.T) {
	var tests = []struct {
		name     string
		expected []int
		round    int
		result   int
	}{
		{
			name:     "Round 1",
			expected: []int{2, 4, 3, 6},
			round:    1,
			result:   20,
		}, {
			name:     "Round 20",
			expected: []int{99, 97, 8, 103},
			round:    20,
			result:   10605,
		}, {
			name:     "Round 1000",
			expected: []int{5204, 4792, 199, 5192},
			round:    1000,
			result:   10605,
		}, {
			name:     "Round 10000",
			expected: []int{52166, 47830, 1938, 52013},
			round:    10000,
			result:   2713310158,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			monkeys, _ := MonkeyInTheMiddle(createMonkeysExample(), tt.round, InspectItemPart2)
			expected := tt.expected
			output := []int{monkeys[0].inspect, monkeys[1].inspect, monkeys[2].inspect, monkeys[3].inspect}
			outputResult := fmt.Sprintf("%v", output)
			expectedResult := fmt.Sprintf("%v", expected)
			if outputResult != expectedResult {
				t.Fatalf("\nexpecetd %s \n got %s", expectedResult, outputResult)
			}

		})
	}

}
