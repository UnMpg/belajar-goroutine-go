package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var Locker = sync.Mutex{}
var Cond = sync.NewCond(&Locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	Cond.L.Lock()
	Cond.Wait()

	fmt.Println("done", value)
	Cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			Cond.Signal()
		}
	}()

	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		time.Sleep(1 * time.Second)
	// 		Cond.Broadcast()
	// 	}

	// }()

	group.Wait()
}
