// ˅
package factory

// ˄

type Link struct {
	// ˅

	// ˄

	ILink

	Name string

	Url string

	// ˅

	// ˄
}

func NewLink(name string, url string) *Link {
	// ˅
	return &Link{Name: name, Url: url}
	// ˄
}

// ˅

// ˄
