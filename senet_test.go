package main

import (
	"reflect"
	"testing"
)

func TestThrowSticks(t *testing.T) {
	var tests = []struct {
		expectedLo int
		expectedHi int
	}{
		{0, 5},
		{0, 5},
		{0, 5},
		{0, 5},
		{0, 5},
	}

	for _, test := range tests {
		output := throwSticks()
		if output < test.expectedLo || output > test.expectedHi {
			t.Error("Test Failed: {} expected low, {} expected high, recieved: {}", test.expectedLo, test.expectedHi, output)
		}
	}
}

func TestValidMove(t *testing.T) {
	var tests = []struct {
		inputTok   int
		inputN     int
		inputBoard Board
		expected   []int
	}{
		{P, 5, Board{
			S, P, S, P, S, P, S, P, S, P,
			S, P, S, P, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[]int{1, 3, 5, 7, 9, 11, 13}},
		{P, 5, Board{
			S, P, S, P, S, P, S, S, P, P,
			S, P, S, P, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[]int{5, 9, 11, 13}},
	}

	for _, test := range tests {
		if output := validMoves(test.inputTok, test.inputN, (test.inputBoard)); !reflect.DeepEqual(output, test.expected) {
			t.Error("Test Failed: {} expected, recieved: {}", test.expected, output)
		}
	}
}

//FIXME add TestUpdateBoard
