package strategy

import (
	"log"
	"time"
)

func BackPressure() {
	data := make(chan int)
	done := make(chan bool)

	// Goroutine pengirim
	go func() {
		for i := 0; ; i++ {
			select {
			case data <- i:
				log.Println("Send data:", i)

			case <-done:
				log.Println("Send data success")
				return
			}
			time.Sleep(time.Second)
		}
	}()

	// Goroutine penerima
	go func() {

		for d := range data {
			time.Sleep(2 * time.Second)
			log.Println("Receive data:", d)
			log.Println("Process data data:", d)
		}
	}()

	time.Sleep(time.Second * 5)
	log.Println("Stop sending data")
	done <- true

}
