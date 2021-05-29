package main

import (
	"fmt"
	//"sync/atomic"
	//"time"
	"os"
	"os/signal"
)

//producer: liên tục tạo ra một số nguyên dựa trên bộ số factor và đưa vào channel
func Producer(factor int, out chan<- int){
	for i := 0; ; i++ {
		out <- i * factor
	}
}

//consumer: liên tục lấy các số từ channel ra để print
func Consumer(in <-chan int){
	for v := range in {
		fmt.Println(v)
	}
}

func main(){
	// tạo hàng đợi
	channel := make(chan int, 64)

	go Producer(3, channel)

	go Producer(5, channel)

	go Consumer(channel)

	//thoát sau khi chạy một khoảng thời gian nhất định
	//time.Sleep(5 * time.Second)

	//Ctrl+C để thoát
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}