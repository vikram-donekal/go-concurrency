package main

import (
	"../src/DTO"
	"../src/Thread1"
	"../src/Thread2"
	"runtime"
	"time"
)

func main() {

	runtime.GOMAXPROCS(4)

	communicationChannel:=make(chan  DTO.MessageDto)

	go Thread1.PushingRandomDto(communicationChannel)
	go Thread2.PullingDtoFromChannel(communicationChannel)

	time.Sleep(5 *  time.Second)

}
