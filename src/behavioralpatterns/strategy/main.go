package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
A game of rock-scissors-paper.
There are two strategies below.

* When winning a game, show the same hand at the next time.
* Calculate a hand from the previous hand stochastically.
*/

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: strategy.exe RandomSeedNumber1 RandomSeedNumber2")
		fmt.Println("Ex.  : strategy.exe 314 15")
	} else {
		var randomSeed1, _ = strconv.ParseInt(os.Args[1], 10, 64)
		var randomSeed2, _ = strconv.ParseInt(os.Args[2], 10, 64)
		player1 := NewPlayer("Emily", NewStrategyA(randomSeed1))
		player2 := NewPlayer("James", NewStrategyB(randomSeed2))

		for i := 0; i < 100; i++ {
			var nextHand1 = player1.NextHand()
			var nextHand2 = player2.NextHand()
			if nextHand1.IsStrongerThan(nextHand2) {
				fmt.Println("Winner: " + player1.ToString())
				player1.Won()
				player2.Lost()
			} else if nextHand2.IsStrongerThan(nextHand1) {
				fmt.Println("Winner: " + player2.ToString())
				player1.Lost()
				player2.Won()
			} else {
				fmt.Println("Draw...")
				player1.Drew()
				player2.Drew()
			}
		}
		fmt.Println("RESULT:")
		fmt.Println(player1.ToString())
		fmt.Println(player2.ToString())
	}
}
