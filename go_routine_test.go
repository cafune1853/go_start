package go_start

import (
	"fmt"
	"testing"
	"time"
)

func TestK(t *testing.T) {
	arr := [4]int{1, 3, 7, 9}
	requesting := make(chan bool, 2)
	for i := range arr {
		requesting <- true
		go func() {
			fmt.Printf("%p %p\n", &i, &arr)
			<-requesting
		}()
	}
	time.Sleep(time.Second)

}

func TestPrintN(t *testing.T) {
	printN(10)
}

func printN(n int) {
	signalArr := make([]chan bool, n+1)
	for i := 0; i <= n; i++ {
		signalArr[i] = make(chan bool, 1)
	}
	signalArr[0] <- true
	for i := 1; i <= n; i++ {
		go printCur(signalArr[i-1], signalArr[i], i)
	}
	<-signalArr[n]
}

// suit
func printCur(preSignal chan bool, curSignal chan bool, cur int) {
	<-preSignal
	fmt.Println(cur)
	curSignal <- true
}

// suit
func TestSumToN(t *testing.T) {
	fmt.Println(sumToN(10))
}
func sumToN(n int) int {
	r := make(chan int, 2)
	go sumToChan(1, n/2, r)
	go sumToChan(n/2+1, n, r)
	return <-r + <-r
}

func sumToChan(s, e int, r chan int) {
	var sum int
	for i := s; i <= e; i++ {
		sum += i
	}
	r <- sum
}
