// ˅
package main

import "strings"

// ˄

type LargeSizeString struct {
	// ˅

	// ˄

	largeSizeChars []*LargeSizeChar

	// ˅

	// ˄
}

func NewLargeSizeString(displayData string) *LargeSizeString {
	// ˅
	largeSizeString := &LargeSizeString{}
	for _, ch := range strings.Split(displayData, "") {
		largeSizeChar := NewLargeSizeCharFactory().GetLargeSizeChar(ch)
		largeSizeString.largeSizeChars = append(largeSizeString.largeSizeChars, largeSizeChar)
	}
	return largeSizeString
	// ˄
}

func (self *LargeSizeString) Display() {
	// ˅
	for i := 0; i < len(self.largeSizeChars); i++ {
		self.largeSizeChars[i].Display()
	}
	// ˄
}

// ˅

// ˄
