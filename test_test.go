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
	fmt.Printf("%q %x %T", []byte("kk"), "中国", "sdfsf")
}
