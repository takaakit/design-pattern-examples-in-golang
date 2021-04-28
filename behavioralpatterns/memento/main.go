package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
Dice game collecting fruits.

* A gamer shakes a dice and the number determine the next state.
* Gamer's money increases or decreases depending on the number. The gamer sometimes gets desserts.
* The game is over if the gamer's money runs out.
*/

func main() {
	gamer := NewGamer(100)
	memento := gamer.CreateMemento()

	for i := 0; i < 100; i++ {
		fmt.Println("==== " + strconv.Itoa(i))
		fmt.Println("Current state: " + gamer.ToString())

		gamer.Play()

		fmt.Println("Gamer's money is " + strconv.Itoa(gamer.money) + ".")

		if gamer.money > memento.money {
			fmt.Println("(Save the current state because money has increased.)")
			memento = gamer.CreateMemento()
		} else if gamer.money < memento.money/2 {
			fmt.Println("(Go back to the previous state because money has decreased.)")
			gamer.RestoreMemento(memento)
		}

		time.Sleep(1 * time.Second)

		fmt.Println("")
	}
}
