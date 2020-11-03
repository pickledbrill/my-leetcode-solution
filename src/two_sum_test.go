package main

import "testing"

func TestTwoSum(t *testing.T) {
	// setup
	input := []int{3, 2, 4}
	target := 6
	expected := []int{1, 2}
	// execute
	result := twoSum(input, target)
	for _, k := range result {
		println(k)
	}
	// assert
	for index, value := range expected {
		a := result[index]
		if value != a {
			t.Errorf("result unmatch")
		}
	}
}
