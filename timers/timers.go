package main

import (
	"fmt"
	"time"
)

func fireTimer(timer *time.Timer) {
	<-timer.C

	fmt.Println("Timer Fired")
}

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	stop2 := timer2.Stop()

	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)

	timer3 := time.NewTimer(2 * time.Second)
	fireTimer(timer3)

	stop := timer3.Stop() // Doenst work that return nill
	if stop {
		fmt.Println("Timer 3 stopped")
	}

}
