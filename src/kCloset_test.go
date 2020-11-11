package main

import "testing"

// [[1,3],[-2,2],[2,-2]]
// 2

func TestKCloset_WithExampleOne_ShouldReturnCorrect(t *testing.T) {
	// setup
	input := [][]int{{1, 3}, {-2, 2}}
	k := 1
	expected := [][]int{{-2, 2}}
	// execute
	result := kClosest(input, k)
	// assert
	if result[0][0] != expected[0][0] || result[0][1] != expected[0][1] {
		t.Errorf("result unmatch")
	}
}

func TestKCloset_WithExampleTwo_ShouldReturnCorrect(t *testing.T) {
	// setup
	input := [][]int{{3, 3}, {5, -1}, {-2, 4}}
	k := 2
	expected := [][]int{{3, 3}, {-2, 4}}
	// execute
	result := kClosest(input, k)
	// assert
	if result[0][0] != expected[0][0] || result[0][1] != expected[0][1] || result[1][0] != expected[1][0] || result[1][1] != expected[1][1] {
		t.Errorf("result unmatch")
	}
}

func TestKCloset_WithExampleExtra_ShouldReturnCorrect(t *testing.T) {
	// setup
	input := [][]int{{1, 3}, {-2, 2}, {2, -2}}
	k := 2
	expected := [][]int{{2, -2}, {-2, 2}}
	// execute
	result := kClosest(input, k)
	// assert
	if result[0][0] != expected[0][0] || result[0][1] != expected[0][1] || result[1][0] != expected[1][0] || result[1][1] != expected[1][1] {
		t.Errorf("result unmatch")
	}
}

func TestKCloset_WithExampleExtraTwo_ShouldReturnCorrect(t *testing.T) {
	// setup
	input := [][]int{{5, -3}, {3, -4}, {-2, 9}, {-8, -8}, {-6, 5}, {6, 1}, {5, 9}}
	k := 6
	expected := [][]int{{3, -4}, {5, -3}, {6, 1}, {-6, 5}, {-2, 9}, {5, 9}}
	// execute
	result := kClosest(input, k)
	// assert
	if result[0][0] != expected[0][0] || result[0][1] != expected[0][1] || result[1][0] != expected[1][0] || result[1][1] != expected[1][1] || result[2][0] != expected[2][0] || result[2][1] != expected[2][1] || result[3][0] != expected[3][0] || result[3][1] != expected[3][1] || result[4][0] != expected[4][0] || result[4][1] != expected[4][1] || result[5][0] != expected[5][0] || result[5][1] != expected[5][1] {
		t.Errorf("result unmatch")
	}
}
