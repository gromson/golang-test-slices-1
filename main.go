package main

import (
	"fmt"
	"sort"
)

func main() {
	res1 := findElements([]int{1, 5, 3, 2, 13, 6, 7, 56, 34, 23, 79})
	res2 := findElements([]int{3, 4, 6, 7, 34, 45, 23, 13})
	fmt.Println(res1)
	fmt.Println(res2)
}

func findElements(src []int) [][3]int {
	var slice sort.IntSlice = src
	slice.Sort()

	var result [][3]int

	for i, v := range slice {
		for j := i + 1; j < len(slice); j++ {
			find := v + slice[j]
			findIn := slice[j+1:]
			if search(find, findIn) {
				r := [3]int{v, slice[j], find}
				result = append(result, r)
			}
		}
	}

	return result
}

func search(val int, in []int) bool {
	if len(in) > 0 {
		index := len(in) / 2
		//fmt.Printf("search %d in %v by index %d\n", val, in, index)

		if val < in[index] {
			return search(val, in[:index])
		} else if val > in[index] {
			return search(val, in[index+1:])
		} else {
			return true
		}
	}

	return false
}
