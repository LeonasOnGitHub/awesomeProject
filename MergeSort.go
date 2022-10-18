package main

import (
	"fmt"
)

func main() {
	//start
	a := [10]int{6, 9, 4, 2, 1, 3, 7, 0, 5, 10}
	fmt.Println("unsorted:\n", a)
	var s []int = a[0:len(a)]
	s = mergeSort(s)
	fmt.Println("sorted:\n", s)
}

func mergeSort(ws []int) []int {
	if len(ws) <= 1 {
		return ws
	}
	var a []int = ws[0 : len(ws)/2]
	a = mergeSort(a)

	var b []int = ws[len(ws)/2 : len(ws)]
	b = mergeSort(b)

	return merge(a, b)
}

func merge(a []int, b []int) []int {

	c := make([]int, len(a)+len(b))
	i := 0
	j := 0

	for k := 0; k < len(c); k++ {
		if i >= len(b) {
			c[k] = a[j]
			j++
		} else if j >= len(a) {
			c[k] = b[i]
			i++
		} else if a[j] > b[i] {
			c[k] = b[i]
			i++
		} else {
			c[k] = a[j]
			j++
		}
	}
	return c
}
