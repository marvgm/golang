package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	
	channel := make(chan int)

	//função anonima que roda 3 workes em multthread: faz o loop de 100 numeros 3 vezes em paralelo
	go func() {
		for i :=1; i <= 30; i++ {
			go worker(channel, i)
		}
	}()
	
   for i :=0; i < 100; i++ {
	   channel <- i
   }
}

func worker(channel chan int, worker int)  {
	for i := range channel  {
		fmt.Println("O numero", i , "processado para o worker:", worker)
		time.Sleep(time.Duration(rand.Intn(4)) * time.Second)
	}
}