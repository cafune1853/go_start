package go_start

import (
	"fmt"
	"testing"
)


// 使用一块连续的一维空间来处理
func allocateSlice(x, y int) [][]int{
	r := make([][]int, y)
	total := make([]int, x * y)
	for i := range r{
		r[i], total = total[:x], total[x:]
	}
	return r
}
// 各自分配， better
func allocateSlice2(x, y int) [][]int{
	r := make([][]int, y)
	for i := range r{
		r[i] = make([]int, x)
	}
	return r
}


func TestTwoDimensionalSliceAllocate(t *testing.T) {
	x := allocateSlice(3, 3)
	for i := range x{
		fmt.Printf("%v %v\n", len(x[i]), cap(x[i]))
	}
	fmt.Println("----------------")
	x = allocateSlice2(3, 3)
	for i := range x{
		fmt.Printf("%v %v\n", len(x[i]), cap(x[i]))
	}
}

