package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var total struct {
	sync.Mutex

	value int
}

var atomicTotal uint64


func worker(wg *sync.WaitGroup) {
	//thông báo hoàn thành khi ra khỏi hàm
	defer wg.Done()

	var i uint64
	for i = 0; i < 100; i++ {
		atomic.AddUint64(&atomicTotal, 1)
	}

	for i := 0; i < 100; i++ {
		//block gorountines khác cho đến khi unlock
		total.Lock()

		//critical section
		total.value++

		//unlock
		total.Unlock()
	}
}

func main(){
	// xử lý các tác vụ atomic (đơn nguyên)

	var wg sync.WaitGroup

	//wg cần chờ 2 gorountines khác
	wg.Add(2)

	go worker(&wg)
	go worker(&wg)

	wg.Wait()

	fmt.Println(atomicTotal)
	fmt.Println(total.value)
}