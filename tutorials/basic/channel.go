package main

import "fmt"

func main(){
	//buffered channel
	done := make(chan int)

	go func(){
		fmt.Println("Hello World")

		// gửi giá trị vào channel thông báo kết thúc gorountine này
		done <- 1
	}()

	//main thread nhận giá trị từ channel và thoát khỏi trạng thái block
	<- done
}