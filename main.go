package main

import (
	"sync"
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
	wg.Add(1)
	go emailWorker(i,recipientChannel,&wg)

	}

	wg.Wait()

}