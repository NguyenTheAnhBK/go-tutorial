package main

import (
	"time"
	"strings"
	"fmt"

	"pubsub/pubsub"
)

func main(){
	//khởi tạo 1 publisher
	p := pubsub.NewPublisher(100 * time.Millisecond, 10)

	//đảm bảo p được đóng trước khi exit
	defer p.Close()

	// `all` subscribe hết các topic
	all := p.Subscribe()

	//subscribe topic chứa golang
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	//publish ra 2 topic
	p.Publish("hello, world!")
	p.Publish("hello, golang!")

	//print những gì subscriber `all` nhận được
	go func(){
		for msg := range all{
			fmt.Println("[all]: ", msg)
		}
	}()

	//print những gì subscribe `golang` nhận được
	go func(){
		for msg := range golang{
			fmt.Println("[golang]: ", msg)
		}
	}()

	//thoát ra sau khi chạy 3s
	time.Sleep(3 * time.Second)
}