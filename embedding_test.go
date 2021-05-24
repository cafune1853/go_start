package go_start

import (
	"fmt"
	"testing"
)

type A struct {
	Q string
}

func (a A) String() string {
	return "A"
}

type B struct {
}

//func (b B) String() string{
//	return "B"
//}

type AB struct {
	*A
	*B
	kk int
}

//func (ab AB) String() string{
//	return fmt.Sprintf("%v", ab.kk)
//}

func TestEmbeddedStruct(t *testing.T) {
	ab := AB{&A{}, &B{}, 10}
	fmt.Println(ab.A.Q)
}
