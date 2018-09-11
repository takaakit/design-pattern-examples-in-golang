// ˅
package main

// ˄

type LargeSizeCharFactory struct {
	// ˅

	// ˄

	poolChars map[string]*LargeSizeChar

	// ˅

	// ˄
}

func NewLargeSizeCharFactory() *LargeSizeCharFactory {
	// ˅
	if instanceLargeSizeCharFactory == nil {
		instanceLargeSizeCharFactory = &LargeSizeCharFactory{}
		instanceLargeSizeCharFactory.poolChars = map[string]*LargeSizeChar{}
	}
	return instanceLargeSizeCharFactory
	// ˄
}

// Create an instance of the large size character.
func (self *LargeSizeCharFactory) GetLargeSizeChar(charName string) *LargeSizeChar {
	// ˅
	lsc := self.poolChars[charName]
	if lsc == nil {
		lsc = NewLargeSizeChar(charName) // Create an instance
		self.poolChars[charName] = lsc
	}
	return lsc
	// ˄
}

// ˅
var instanceLargeSizeCharFactory *LargeSizeCharFactory = nil

// ˄
