package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	var mutex sync.Mutex
	x := 0

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("couter = ", x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balace  int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balace = account.Balace + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalace() int {
	account.RWMutex.RLock()
	belace := account.Balace
	account.RWMutex.RUnlock()

	return belace
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalace())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("total Balace", account.GetBalace())
}

type UserBalace struct {
	sync.Mutex
	Name   string
	Balace int
}

func (user *UserBalace) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalace) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalace) Change(amount int) {
	user.Balace = user.Balace + amount
}

func Transfer(user1 *UserBalace, user2 *UserBalace, amuount int) {
	user1.Lock()
	fmt.Println("Lock User1", user1.Name)
	user1.Change(-amuount)

	time.Sleep(1 * time.Second)
	user2.Lock()
	fmt.Println("Lock user 2", user2.Name)
	user2.Change(amuount)

	time.Sleep(1 * time.Second)
	user1.Unlock()
	user2.Unlock()
}

func TestDeadLock(t *testing.T) {
	user1 := UserBalace{
		Name:   "fendy",
		Balace: 1000,
	}

	user2 := UserBalace{
		Name:   "Agus",
		Balace: 2000,
	}

	go Transfer(&user1, &user2, 100)
	go Transfer(&user2, &user1, 200)

	time.Sleep(2 * time.Second)

	fmt.Println("User ", user1.Name, ", Balace ", user1.Balace)
	fmt.Println("User ", user2.Name, ", Balace ", user2.Balace)
}
