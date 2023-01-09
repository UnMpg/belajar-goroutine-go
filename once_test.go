package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
)

var Counter = 0

func OnlyOnce() {
	Counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			once.Do(OnlyOnce)
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println("conternya", Counter)

}
