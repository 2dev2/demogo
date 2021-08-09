package main

import (
	"context"
	"fmt"
	consumer2 "github.com/2dev2/demogo/consumer"
	"github.com/2dev2/demogo/producer"
	parser2 "github.com/2dev2/demogo/parser"
	"github.com/2dev2/demogo/reader"
	"time"

	"sync"
)

func main(){

	inputWord:="RoboEarth"
	fmt.Print("main")

	result:=[]string{}
	filereader:=reader.NewFileReader()
	urls:=make(chan string)

	words:=make(chan parser2.TupleMap,100)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		filereader.Read("", urls)
	}()



	parser:=parser2.NewWhileSpaceSeparatorParser(words)
	producer:= producer.NewProducer(urls,parser)
	consumer:= consumer2.NewConsumer(words,urls)
	dict:=map[string][]string{}
	var wg sync.WaitGroup
	producerDone:=make(chan int)

	var mutex = &sync.Mutex{}
	go func() {
		for{
			select {
			case url:=<-urls:
				producer.Produce(url)
			default:
				producerDone<-1
			}
		}
	}()
	wg.Add(1)
	go consumer.Worker(ctx,&wg,dict,mutex)
	<-producerDone
	time.Sleep(200*time.Second) //this is hard limit for now we will check the result later
	fmt.Print("============ producerDone Done========")
	cancel()
	//producer.Close()
	wg.Wait()
	fmt.Print("============ consumer Done========")
	fmt.Print("** result **")

	//handle the multi separated word by While space separator them and put count of math Then intersection the  two list
	for k,v:=range dict{
		if k==inputWord{
			result = v
		}
	}

	fmt.Print(result)


}


