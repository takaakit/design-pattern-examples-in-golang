// ˅
package main

import (
	"bufio"
	"fmt"
	"os"
)

// ˄

type LargeSizeChar struct {
	// ˅

	// ˄

	charName string

	// Display data of the large size character
	displayData string

	// ˅

	// ˄
}

func NewLargeSizeChar(charName string) *LargeSizeChar {
	// ˅
	largeSizeChar := &LargeSizeChar{}
	largeSizeChar.charName = charName

	file, err := os.Open("big" + charName + ".txt")
	if err != nil {
		fmt.Println(err)
		largeSizeChar.displayData = charName + "?"
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			largeSizeChar.displayData += scanner.Text() + "\n"
		}
	}

	return largeSizeChar
	// ˄
}

// Display the large size character
func (self *LargeSizeChar) Display() {
	// ˅
	fmt.Println(self.displayData)
	// ˄
}

// ˅

// ˄
