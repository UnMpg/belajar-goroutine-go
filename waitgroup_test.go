package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsychronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("hallo")
	time.Sleep(1 * time.Second)
}

func TestWaitgroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsychronous(group)
	}

	group.Wait()
	fmt.Println("complete")
}
