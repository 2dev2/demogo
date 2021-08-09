package main

import (
	"context"
	"fmt"
	consumer2 "github.com/2dev2/demogo/consumer"
	parser2 "github.com/2dev2/demogo/parser"
	"github.com/2dev2/demogo/producer"
	"github.com/2dev2/demogo/reader"
	"sync"
	"time"
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


	dict:=map[string][]string{}
	parser:=parser2.NewWhileSpaceSeparatorParser(words)
	producer:= producer.NewProducer(urls,parser)
	consumer:= consumer2.NewConsumer(words,urls,dict)

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
	go consumer.Worker(ctx,&wg,mutex)
	<-producerDone
	go func(){
		//wait max 20 Seconds we should be configuring this
		time.Sleep(20*time.Second)
		cancel()
	}()
	//time.Sleep(200*time.Second) //this is hard limit for now we will check the result later
	fmt.Print("============ producerDone Done========")

	//producer.Close()
	wg.Wait()
	fmt.Print("============ consumer Done========")
	fmt.Print("   ** result **    ")

	//handle the multi separated word by While space separator them and put count of math Then intersection the  two list
	for k,v:=range dict{
		if k==inputWord{
			result = v
		}
	}

	fmt.Print(result)


}


