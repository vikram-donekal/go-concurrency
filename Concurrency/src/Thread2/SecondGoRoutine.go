package Thread2

import ("log"
 "../DTO"
)

func  PullingDtoFromChannel( inputChannel <-chan DTO.MessageDto )  {

	for {
		if msgInput,ok:= <- inputChannel;ok {
			log.Println(msgInput)
		}else {
			break
		}

	}
}
