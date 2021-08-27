// ˅
package memento

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ˄

type Gamer struct {
	// ˅

	// ˄

	// Gamer's money
	money int

	// ˅

	// ˄
}

func NewGamer(money int) *Gamer {
	// ˅
	rand.Seed(time.Now().UnixNano())

	return &Gamer{money: money}
	// ˄
}

func (g *Gamer) CreateMemento() *Memento {
	// ˅
	return NewMemento(g.money)
	// ˄
}

func (g *Gamer) SetMemento(memento *Memento) {
	// ˅
	g.money = memento.Money()
	// ˄
}

// Play a game
func (g *Gamer) Play() {
	// ˅
	dice := rand.Intn(6) + 1
	fmt.Println("The number of dice is " + strconv.Itoa(dice) + ".")

	preMoney := g.money
	switch dice {
	case 1, 3, 5:
		// In case of odd...Money is halved
		g.money /= 2
		fmt.Println("Gamer's money is halved: " + strconv.Itoa(preMoney) + " -> " + strconv.Itoa(g.money))
	case 2, 4, 6:
		// In case of even...Money doubles
		g.money *= 2
		fmt.Println("Gamer's money doubles: " + strconv.Itoa(preMoney) + " -> " + strconv.Itoa(g.money))
	default:
		// Other...Exit
		fmt.Println("Unexpected value.")
		os.Exit(1)
	}
	// ˄
}

func (g *Gamer) Money() int {
	// ˅
	return g.money
	// ˄
}

func (g *Gamer) String() string {
	// ˅
	return "[money = " + strconv.Itoa(g.money) + "]"
	// ˄
}

// ˅

// ˄
