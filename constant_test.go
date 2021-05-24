package go_start

import (
	"testing"
)

type ByteSize float64

const (
	_ = iota
	KB
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func TestConstant(t *testing.T) {
	println(KB, MB, GB)
}
