package main

import "fmt"

func main() {
func main() {
	// init
	var buf0 []int             // len=0, cap=0
	buf1 := []int{}            // len=0, cap=0
	buf2 := []int{42}          // len=1, cap=1
	buf3 := make([]int, 0)     // len=0, cap=0
	buf4 := make([]int, 5)     // len=5, cap=5
	buf5 := make([]int, 5, 10) // len=5, cap=10

	println(buf0, buf1, buf2, buf3, buf4, buf5)

	// get value at index
	someInt := buf2[0]

	// run time error:

	// panic: runtime error: index out of range
	// someOtherInt := buf2[1]

	fmt.Println(someInt)

	// adds items
	var buf []int            // len=0, cap=0
	buf = append(buf, 9, 10) // len=2, cap=2
	buf = append(buf, 12)    // len=3, cap=4

	// adds other slice
	otherBuf := make([]int, 3)     // [0,0,0]
	buf = append(buf, otherBuf...) // len=6, cap=8

	fmt.Println(buf, otherBuf)

	// slice info
	var bufLen, bufCap int = len(buf), cap(buf)

	fmt.Println(bufLen, bufCap)
}
