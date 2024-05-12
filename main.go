package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wait_group sync.WaitGroup // declaração
	wait_group.Add(2)             // quantidade de goroutines que vão estar rodando ao mesmo tempo /2 goroutines do grupo deespera
	go func() {
		writePhrase("Hello world")
		wait_group.Done()
	}()
	go func() {
		writePhrase("Olá mundo")
		wait_group.Done()
	}()

	wait_group.Wait()
}

func writePhrase(text string) {
	for k := 0; k < 5; k++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
