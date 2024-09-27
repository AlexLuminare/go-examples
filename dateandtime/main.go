package main

import (
	//"fmt"
	"time"
)

func writeToChannel(channel chan<- string) {
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	for _, name := range names {
		channel <- name
		//time.Sleep(time.Second * 1)
	}
	close(channel)
}
func main() {
	nameChannel := make(chan string)
	time.AfterFunc(time.Second*5, func() {
		writeToChannel(nameChannel)
	})
	for name := range nameChannel {
		Printfln("Read name: %v", name)
	}
}
