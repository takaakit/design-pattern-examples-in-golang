// ˅
package factory

// ˄

type Data struct {
	// ˅

	// ˄

	IData

	Name string

	Items []Item

	// ˅

	// ˄
}

func NewData(name string) *Data {
	// ˅
	return &Data{Name: name, Items: []Item{}}
	// ˄
}

func (d *Data) Add(item Item) {
	// ˅
	d.Items = append(d.Items, item)
	// ˄
}

// ˅

// ˄
