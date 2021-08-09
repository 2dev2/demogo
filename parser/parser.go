package parser

import (
	"strings"
)


type TupleMap struct{
	Word string
	Url string
}
type ParserContent interface {
	Parse(content string,url string) error
	Close()
}
type whileSpaceSeparatorParser struct{
	producer chan TupleMap
}


//we can create seperate parse based on anchor tags
func NewWhileSpaceSeparatorParser(producer chan TupleMap) *whileSpaceSeparatorParser{
	 return &whileSpaceSeparatorParser{
		 producer:producer,
	 }
}

func(p *whileSpaceSeparatorParser) Parse(content string,url string) error{

	listOfWord:= strings.Split(content," ")
	for _,v:=range listOfWord{
		p.producer<-TupleMap{v,url}
	}
	return nil

}


func(p *whileSpaceSeparatorParser) Close(){
	close(p.producer)
}