package memento

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

/*
A dice game in which money increases and decreases:
* A gamer shakes a dice and the number determine the next state.
* If the number of dice is even, gamer's money doubles, and if it is odd, gamer's money is halved.
* If the gamer's money is less than half of the highest amount, it returns to the highest amount.
* The game is repeated.
*/

func TestMain(m *testing.M) {
	gamer := NewGamer(100)
	memento := gamer.CreateMemento()

	for i := 0; i < 10; i++ {
		fmt.Println("==== Turn " + strconv.Itoa(i+1)) // Display count

		gamer.Play() // Play a game

		// Determine the behavior of the Memento
		if gamer.Money() > memento.Money() {
			fmt.Println("(Gamers' money is the highest ever, so record the current state.)")
			memento = gamer.CreateMemento()
		} else if gamer.Money() < memento.Money()/2 {
			fmt.Println("(Gamer's money is less than half of the highest amount, so return to the recorded state.)")
			gamer.SetMemento(memento)
			fmt.Println("Gamer's money returns to " + strconv.Itoa(gamer.Money()) + ".")
		}

		fmt.Println()

		time.Sleep(1 * time.Second)
	}
}
