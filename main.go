package main

import (
	"sync"
	"time"
)



type Recipient struct{
	Name string
	Email string
}


func main(){

	recipientChannel:=make(chan Recipient)

	go func(){
		loadRecipient("./emails.csv",recipientChannel)


	}()


	var wg sync.WaitGroup

	workerCnt:=5
	for i:=1; i<=workerCnt;i++{
	go emailWorker(i,recipientChannel,&wg)

	}

	time.Sleep(3*time.Second)
}