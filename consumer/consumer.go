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
	//mu sync.Mutex
}

func NewConsumer(words chan parser2.TupleMap,urls chan string) *consumer{
	return &consumer{words,urls}
}

func(c *consumer) Worker(ctx context.Context, wg *sync.WaitGroup,dict map[string][]string,mu *sync.Mutex) {
	defer func() {
		fmt.Print("============ Worker done========")
		wg.Done()
	}()
	for{
		select {
			case v:=<-c.words:
				mu.Lock()
				dict[v.Word] = append(dict[v.Word], v.Url)
				mu.Unlock()
			case <-ctx.Done():
				return
		}
	}
}
