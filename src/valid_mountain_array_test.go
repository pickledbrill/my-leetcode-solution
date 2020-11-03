package main

import "testing"

func TestValidMountainArray_WithLongArray_ShouldReturnSuccess(t *testing.T) {
	// setup
	input := []int{0, 2, 3, 4, 5, 2, 1, 0}
	expected := true
	// execute
	result := validMountainArray(input)
	// assert
	if result != expected {
		t.Errorf("result unmatch")
	}
}

func TestValidMountainArray_WithLongArray_ShouldReturnFail(t *testing.T) {
	// setup
	input := []int{1, 7, 9, 5, 4, 1, 2}
	expected := false
	// execute
	result := validMountainArray(input)
	// assert
	if result != expected {
		t.Errorf("result unmatch")
	}
}

func TestValidMountainArray_WithShortArray_ShouldReturnSuccess(t *testing.T) {
	// setup
	input := []int{0, 3, 2, 1}
	expected := true
	// execute
	result := validMountainArray(input)
	// assert
	if result != expected {
		t.Errorf("result unmatch")
	}
}

func TestValidMountainArray_WithShortArray_ShouldReturnFail(t *testing.T) {
	// setup
	input := []int{0, 3, 2, 5}
	expected := false
	// execute
	result := validMountainArray(input)
	// assert
	if result != expected {
		t.Errorf("result unmatch")
	}
}
