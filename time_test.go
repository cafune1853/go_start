package go_start

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	tt := time.Now()
	tt = tt.AddDate(0, 0, 0)
	tt = tt.AddDate(0, 0, -1)
	tt, err := time.Parse("2006-01-02", "2021-02-28")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tt)
}
