package producer

import (
	"fmt"
	"github.com/2dev2/demogo/parser"
	"io"
	"net/http"
	"strings"
)

type producer struct{
	Urls chan string
	urlsMap map[string]bool
	parser parser.ParserContent
	//done chan int
}

func NewProducer(urls chan string ,parser parser.ParserContent)*producer{
	return &producer{
		Urls: urls,
		urlsMap: map[string]bool{},
		parser:parser,
	}
}

func(c *producer) Produce(url string){
	//check coorect logic based on <anchor tags>
	if strings.Contains(url,"http") {
		//we can append count/depth of parsing traversal in url for checking purpose
		if _, ok := c.urlsMap[url]; ok {
			return;
		}
		c.urlsMap[url] = true
		//put some circuit break logic
		//rate limit  by using semaphore/buffer channel
		//retry http.do support that
		resp, err := http.Get(string(url))
		if err != nil {
			fmt.Print("got error when fetching", url)
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Print("got error when fetching")
			return
		}
		c.parser.Parse(string(body),url)
	}
	//else{
	//	c.parser.Parse(string(url),url)
	//}

}

func(c *producer) Close(){
	c.parser.Close()
}