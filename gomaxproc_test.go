package golanggoroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxproc(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 200; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCPU := runtime.NumCPU()
	fmt.Println("total CPU ", totalCPU)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total thread ", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("total Goroutine ", totalGoroutine)

	group.Wait()
}

func TestChangeThread(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 200; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCPU := runtime.NumCPU()
	fmt.Println("total CPU ", totalCPU)

	runtime.GOMAXPROCS(10)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total thread ", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("total Goroutine ", totalGoroutine)

	group.Wait()
}
