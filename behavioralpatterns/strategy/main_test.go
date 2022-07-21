package strategy

import (
	"fmt"
	"testing"
)

/*
A game of rock-scissors-paper. Two strategies are available:
* Random Strategy: showing a random hand signal.
* Mirror Strategy: showing a hand signal from the previous opponent's hand signal.
*/

func TestMain(m *testing.M) {
	player1 := NewPlayer("Emily", NewRandomStrategy())
	player2 := NewPlayer("James", NewMirrorStrategy())

	for i := 0; i < 100; i++ {
		handOfPlayer1 := player1.ShowHandSignal()
		handOfPlayer2 := player2.ShowHandSignal()

		// Judge win, loss, or draw
		var resultOfPlayer1 GameResultType
		var resultOfPlayer2 GameResultType
		if handOfPlayer1.IsStrongerThan(handOfPlayer2) {
			fmt.Println("Winner: " + player1.ToString())
			resultOfPlayer1 = Win
			resultOfPlayer2 = Loss
		} else if handOfPlayer2.IsStrongerThan(handOfPlayer1) {
			fmt.Println("Winner: " + player2.ToString())
			resultOfPlayer1 = Loss
			resultOfPlayer2 = Win
		} else {
			fmt.Println("Draw...")
			resultOfPlayer1 = Draw
			resultOfPlayer2 = Draw
		}

		player1.NotifyGameResult(resultOfPlayer1, handOfPlayer1, handOfPlayer2)
		player2.NotifyGameResult(resultOfPlayer2, handOfPlayer2, handOfPlayer1)
	}
	fmt.Println("RESULT:")
	fmt.Println(player1.ToString())
	fmt.Println(player2.ToString())
}
