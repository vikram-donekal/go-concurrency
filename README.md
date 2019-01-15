# go-concurrency

This example in go lang about concurrency. 

Description:
	Sample application which will showcase concurrency in go lang using go-routines.
	
	channel is been used as communication between two go-routines.
	
	I made as simple as possible to stimulate java Threading concepts.
	
	
How it works:
	Two go-routines which will communicate over channel.
	data exchanged over a channel between two go-routines is a randam string.
	
	

How to start:
	1:Using Dokcer 
		$ docker run -d vikramdonekal/go-concurrency
		
		grab logs to see data tranfering between two go routine.
		
		$ docker logs containerId
		
		
