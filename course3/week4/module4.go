package main

import (
	"fmt"
	"sync"
)

func main() {
	chopsticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = &ChopStick{id: i}
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{i, chopsticks[i], chopsticks[(i+1)%5], 3}
	}

	var wg sync.WaitGroup
	host := &Host{c: make(chan int, 2)}
	wg.Add(15)

	for i := 0; i < len(philosophers); i++ {
		go host.givePermission(philosophers[i].id)
		x := <-host.c
		go philosophers[x].eat(&wg)
	}
	wg.Wait()
}

type Host struct {
	c chan int
}

func (h Host) givePermission(philo1 int) {
	h.c <- philo1
}

type ChopStick struct {
	id int
	sync.Mutex
}

type Philosopher struct {
	id              int
	leftCS, rightCS *ChopStick
	hungry          int
}

func (p Philosopher) eat(wg *sync.WaitGroup) {
	for p.hungry > 0 {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("Philosopher %d is eating\n", p.id)
		fmt.Printf("Philosopher %d is done eating\n\n", p.id)
		p.hungry = p.hungry - 1

		p.leftCS.Unlock()
		p.rightCS.Unlock()

		wg.Done()
	}
}
