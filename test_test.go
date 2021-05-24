package go_start

import (
	"fmt"
	"testing"
)

func TestCommon(t *testing.T) {
	type TX struct {
		a int
		b float64
		c string
	}
	tx := &TX{7, -2.35, "abc\tdef"}
	fmt.Printf("%v\n", tx)
	fmt.Printf("%+v\n", tx)
	fmt.Printf("%#v\n", tx)
	fmt.Printf("%q %x %T\n", []byte("kk"), "中国", "sdfsf")

	sl := []int{1, 2, 3}
	pr(sl...)
}

func pr(qq ...int) {
	fmt.Printf("%T\n", qq)
	for x := range qq {
		fmt.Println(x)
	}
}
