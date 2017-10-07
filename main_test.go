package main

import (
	"reflect"
	"testing"
)

var testSrc [][]int = [][]int{
	[]int{1, 5, 3, 2, 8},
	[]int{3, 4, 6, 7, 8, 10},
}

var testRes [][][3]int = [][][3]int{
	[][3]int{{1, 2, 3}, {2, 3, 5}, {3, 5, 8}},
	[][3]int{{3, 4, 7}, {3, 7, 10}, {4, 6, 10}},
}

func TestFindElements(t *testing.T) {
	for i, sl := range testSrc {
		res := findElements(sl)
		if !reflect.DeepEqual(res, testRes[i]) {
			t.Errorf("For %v\nexpected %v\ngot %v", sl, testRes[i], res)
		}
	}
}
