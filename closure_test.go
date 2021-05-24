package go_start

import (
	"fmt"
	"testing"
)

func initCounter() (func(i int), func(i int), func() int) {
	i := 0
	return func(j int) {
			i += j
		}, func(j int) {
			i -= j
		}, func() int {
			return i
		}
}

func TestClosure(t *testing.T) {
	adder, minus, getter := initCounter()
	adder(10)
	minus(3)
	fmt.Println(getter())
}
