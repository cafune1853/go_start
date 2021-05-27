package go_start

import (
	"fmt"
	"github.com/cafune1853/go_start/web/core"
	"testing"
)

func TestCommon(t *testing.T) {
	fmt.Println("kk")
	fmt.Println(core.Page{})
}

type X struct {
	Name string
}

func (x *X)getName()string  {
	return x.Name
}

type Y struct {
	X
	Name string
}
