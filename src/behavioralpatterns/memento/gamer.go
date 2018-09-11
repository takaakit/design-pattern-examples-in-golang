// ˅
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// ˄

type Gamer struct {
	// ˅

	// ˄

	// Gamer's money
	money int

	// Acquired desserts
	desserts []string

	// Dessert name table
	dessertsName []string

	// ˅

	// ˄
}

func NewGamer(money int) *Gamer {
	// ˅
	rand.Seed(time.Now().UnixNano())

	gamer := &Gamer{}
	gamer.money = money
	gamer.dessertsName = []string{"Cake", "Candy", "Cookie"}

	return gamer
	// ˄
}

// Get a dessert
func (self *Gamer) GetDessert() string {
	// ˅
	var prefix = ""
	if rand.Intn(2) == 0 {
		prefix = "Delicious "
	}
	return prefix + self.dessertsName[rand.Intn(3)]
	// ˄
}

// Get current status
func (self *Gamer) CreateMemento() *Memento {
	// ˅
	memento := NewMemento(self.money)
	for i := 0; i < len(self.desserts); i++ {
		var dessert = self.desserts[i]
		if strings.Index(dessert, "Delicious ") == 0 { // Add a only delicious dessert
			memento.addDessert(dessert)
		}
	}
	return memento
	// ˄
}

// Undo status
func (self *Gamer) RestoreMemento(memento *Memento) {
	// ˅
	self.money = memento.money
	self.desserts = memento.desserts
	// ˄
}

// Play a game
func (self *Gamer) Play() {
	// ˅
	var dice = rand.Intn(6) + 1
	switch dice {
	case 1:
		// In case of 1...Gamer's money increases
		self.money += 100
		fmt.Println("Gamer's money increases.")
	case 2:
		// In case of 2...Gamer's money halves
		self.money /= 2
		fmt.Println("Gamer's money halves.")
	case 6:
		// In case of 6...Gamer gets desserts
		var gotDessert = self.GetDessert()
		fmt.Println("Gamer gets desserts(" + gotDessert + ")")
		self.desserts = append(self.desserts, gotDessert)
	default:
		// Other...Nothing happens
		fmt.Println("Nothing happens.")
	}
	// ˄
}

func (self *Gamer) ToString() string {
	// ˅
	var strDesserts string
	for i := 0; i < len(self.desserts); i++ {
		strDesserts += self.desserts[i]
		if i < len(self.desserts)-1 {
			strDesserts += ", "
		}
	}
	return "[money = " + strconv.Itoa(self.money) + ", desserts = " + strDesserts + "]"
	// ˄
}

// ˅

// ˄
