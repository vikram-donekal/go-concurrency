package Thread1

import (
	"../DTO"
	"../Utill"
	"math/rand"
)

var counterInputValue uint64
func PushingRandomDto( inputChannel chan  <- DTO.MessageDto){

	for {

		inputChannel <-  GenerateDto ()
	}


}

func GenerateDto () DTO.MessageDto{
	inputString:=Utill.StringGenerator(rand.Intn(200 - 150) + 15)

	counterInputValue=counterInputValue+1
	inputDto:=DTO.MessageDto{CounterValue : counterInputValue ,Message:inputString}
	return  inputDto
}