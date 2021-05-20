package go_start

import (
	"fmt"
	"testing"
)

type Person struct{
	a int32
	c *byte
	b int32
}



func TestCommon(t *testing.T) {
	var si []int
	var sip = &si
	*sip = make([]int, 12)
	fmt.Printf("%p %v %v", sip, si, si == nil)
}
