// ˅
package flyweight

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
		instanceLargeSizeCharFactory = &LargeSizeCharFactory{poolChars: map[string]*LargeSizeChar{}}
	}
	return instanceLargeSizeCharFactory
	// ˄
}

// Create an instance of the large size character.
func (l *LargeSizeCharFactory) GetLargeSizeChar(charName string) *LargeSizeChar {
	// ˅
	lsc := l.poolChars[charName]
	if lsc == nil {
		lsc = NewLargeSizeChar(charName) // Create an instance
		l.poolChars[charName] = lsc
	}
	return lsc
	// ˄
}

// ˅
var instanceLargeSizeCharFactory *LargeSizeCharFactory = nil

// ˄
