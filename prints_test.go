package go_start

import (
	"fmt"
	"testing"
)

func TestPrints(t *testing.T) {
	// Sprint* Fprint* Print*
	// %d %x for int
	fmt.Printf("%d %x\n", 10, 10)
	// %x for string
	fmt.Printf("%x\n", "中国")
	// %q for quota string/byte
	fmt.Printf("%q %q\n", []byte("A"), "B")
	// %v for common use, %#v for full path and quota string
	fmt.Printf("%#v\n", "xxx")
	// %+v for print key in map and slice index
	p := Person{Name: "Lili", Age: 24}
	fmt.Printf("%+#v %+#v\n", p, []string{"a", "b"})
	// %p for pointer address %T for type
	fmt.Printf("%p %T\n", &p, &p)
}
