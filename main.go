package main

import (
	"fmt"
	"sort"
)

func itemExists(slice []int, s int) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}

func AppendSlice(slice1 []int, slice2 []int) []int {
	var slice []int = slice1
	for _,s := range slice2 {
		if itemExists(slice, s) == false {
			slice = append(slice, s)
		}
	}
	return slice
}

func main() {
	slice1 := []int{4, 2, 1, 3}
	slice2 := []int{3, 6, 5, 4, 3}

	slice := AppendSlice(slice1, slice2)
	sort.Ints(slice)
	fmt.Println(slice)
}