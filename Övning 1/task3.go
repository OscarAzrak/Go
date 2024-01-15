package main

import (
	"fmt"
	"time"
)

func Remind() {

	//NewTicker "ticks" in the background with the selected duration

	//using Seconds for demonstration, change to time.Hour otherwise
	eatTimer := time.NewTicker(3 * time.Second)

	workTimer := time.NewTicker(8 * time.Second)
	sleepTimer := time.NewTicker(24 * time.Second)

	//Using select to identify the right type of timer/case and return the right string
	for {
		select {
		//case works like an if-else statement but the order doesn't matter
		case eat := <-eatTimer.C:
			fmt.Print("Klockan är: ", eat.Hour(), ".", eat.Minute(), ".", eat.Second(), ": ", "Dags att äta", "\n")
		case work := <-workTimer.C:
			fmt.Print("Klockan är: ", work.Hour(), ".", work.Minute(), ".", work.Second(), ": ", "Dags att arbeta", "\n")
		case sleep := <-sleepTimer.C:
			fmt.Print("Klockan är: ", sleep.Hour(), ".", sleep.Minute(), ".", sleep.Second(), ": ", "Dags att sova", "\n")
		}
	}
}

func main() {
	currTime := time.Now()
	fmt.Println("Start time: ", currTime.Format("2006-01-02 15:04:05"))
	Remind()

}
