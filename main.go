package main



type Recipient struct{
	Name string
	Email string
}


func main(){

	recipientChannel:=make(chan Recipient)
	go emailWorker(1,recipientChannel)

	go func(){
		loadRecipient("./emails.csv",recipientChannel)


	}()

}