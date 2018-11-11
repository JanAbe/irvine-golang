package main

import (
	"fmt"
	"sync"
)

func main() {
	// make chopsticks
	chopsticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = &ChopStick{id: i}
	}

	// make philosophers
	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{i, chopsticks[i], chopsticks[(i+1)%5], 3}
	}

	var wg sync.WaitGroup
	host := &Host{c: make(chan int, 2)}
	wg.Add(15)

	// give permission to a philosopher -> pass it into a channel
	// read from the channel -> this philosopher can eat
	// philosophers[x] works because the id and index correspond (this is kinda sketchy though)
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
	// why does it sometimes happen that a philosopher eats more than 3 times?
	// the p.hungry is inside a lock() right?
	// changing it to a normal forloop fixes it
	// for p.hungry > 0 {
	for i := 0; i < 3; i++ {
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
