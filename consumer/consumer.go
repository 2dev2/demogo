package consumer

import (
	"context"
	"fmt"
	parser2 "github.com/2dev2/demogo/parser"
	"sync"
)

type consumer struct{
	words chan parser2.TupleMap
	urls chan string
	dict map[string][]string
	//mu sync.Mutex
}

func NewConsumer(words chan parser2.TupleMap,urls chan string,dict map[string][]string) *consumer{
	return &consumer{words,urls,dict}
}

func(c *consumer) Worker(ctx context.Context, wg *sync.WaitGroup,mu *sync.Mutex) {
	defer func() {
		fmt.Print("============ Worker done========")
		wg.Done()
	}()
	for{
		select {
			case v:=<-c.words:
				mu.Lock()
				c.dict[v.Word] = append(c.dict[v.Word], v.Url)
				mu.Unlock()
			case <-ctx.Done():
				return
		}
	}
}
