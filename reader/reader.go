package reader


type Reader interface{
	Read(path string,urls chan string)
}


type fileReader struct{

}

func NewFileReader() *fileReader{
	return &fileReader{}
}

func (f *fileReader) Read(path string, urls chan string){
	list:=[]string{
		"https://userdocs.rapyuta.io/1_understanding-rio/11_introduction-to-rapyuta-io/",
		//"https://userdocs.rapyuta.io/1_understanding-rio/12_core-concepts/",
	}
	//list:=[]string{
	//	"RoboEarth",
	//}
	for _,v:=range list{
		urls<-v
	}
}