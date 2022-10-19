package main

import (
	"fmt"
)

var countInv = 0

func main() {
	//start
	q1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}                                       // 0 inversions
	q2 := []int{8, 7, 6, 5, 4, 3, 2, 1}                                             // 28
	q3 := []int{54044, 14108, 79294, 29649, 25260, 60660, 2995, 53777, 49689, 9083} // 28 inversions
	q4 := []int{0, 1, 12, 11, 4, 5, 6, 10, 8, 9}                                    // 15 inversions
	fmt.Println("unsorted:\n", q1, "\n", q2, "\n", q3, "\n", q4, "\n")

	var s []int = q1[0:len(q1)]
	s = sortAndCountInv(s)
	fmt.Println("sorted:\n", s, "inv: ", countInv)
	s = q2[0:len(q2)]
	s = sortAndCountInv(s)
	fmt.Println(s, "inv: ", countInv)
	s = q3[0:len(q3)]
	s = sortAndCountInv(s)
	fmt.Println(s, "inv: ", countInv)
	s = q4[0:len(q4)]
	s = sortAndCountInv(s)
	fmt.Println(s, "inv: ", countInv)
}

func sortAndCountInv(ws []int) []int {
	if len(ws) <= 1 {
		return ws
	}
	var a []int = ws[0 : len(ws)/2]
	a = sortAndCountInv(a)

	var b []int = ws[len(ws)/2 : len(ws)]
	b = sortAndCountInv(b)

	return mergeInv(a, b)
}

func mergeInv(a []int, b []int) []int {

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
			countInv++
		} else {
			c[k] = a[j]
			j++
		}
	}
	return c
}

func mCountInv(a []int) int {
	numInv := 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i] > a[j] {
				numInv++
			}
		}
	}
	return numInv
}

func countSplitInv(a []int, b []int) int {
	numInv := 0

	return numInv
}
