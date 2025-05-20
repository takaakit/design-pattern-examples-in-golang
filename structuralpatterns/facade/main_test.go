package facade

import (
	"testing"
)

/*
Create a simple homepage through a Facade (PageCreator). The Facade gets info from
the DataLibrary and uses the info to create an HTML file.
*/

func Test(t *testing.T) {
	pageCreator := NewPageCreator()
	pageCreator.CreateSimpleHomepage("emily@example.com", "Homepage.html")
}
