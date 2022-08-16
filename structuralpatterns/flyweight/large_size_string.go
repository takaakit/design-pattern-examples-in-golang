// ˅
package flyweight

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

func (l *LargeSizeString) Display() {
	// ˅
	for _, largeSizeChar := range l.largeSizeChars {
		largeSizeChar.Display()
	}
	// ˄
}

// ˅

// ˄
