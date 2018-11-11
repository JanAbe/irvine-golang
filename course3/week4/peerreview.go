package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	leftCS, rightCS *ChopS
	id              int
}

func (p *Philo) eat(host chan int) {
	for i := 0; i < 3; i++ {
		//get permission
		host <- 1
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("starting to eat", p.id)
		fmt.Println("finishing eating", p.id)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
		<-host
	}
	wg.Done()
}

func main() {
	//as the host channel has buffer of 2, only 2 philos will dine at any given instance
	host := make(chan int, 2)
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5], i + 1}
	}

	//eat
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat(host)
	}
	wg.Wait()
	//fmt.Scanln()

}
