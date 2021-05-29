package main

import (
	"fmt"
	"sync"
)

func main(){
	var mutex sync.Mutex

	mutex.Lock()

	go func(){
		fmt.Println("Hello World")
		
		mutex.Unlock()
	}()

	mutex.Lock()
}