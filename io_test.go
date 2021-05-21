package go_start

import (
	"fmt"
	"os"
	"testing"
)

func TestIO(t *testing.T) {
	content, err := os.ReadFile("/tmp/hzw_test")
	if err == nil {
		str := string(content)
		fmt.Println(str)
	} else {
		fmt.Println("File not found")
	}
}
