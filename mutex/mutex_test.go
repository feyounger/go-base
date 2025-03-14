package mutex

import (
	"fmt"
	"testing"
	"time"
)

func TestCh(t *testing.T) {
	s1 := make(chan int, 10)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			s1 <- 1
		}
	}()

	i := 1

	for ; i < 4; i++ {
		fmt.Printf("=== times := %d", i)
		select {
		case _ = <-s1:
			if i == 2 {
				fmt.Println(" break")
				break
			} else {
				fmt.Println(" not break")
			}
		}
	}
}
