package abstractfactory

import (
	"fmt"
	"os"

	. "github.com/takaakit/design-pattern-examples-in-golang/creationalpatterns/abstractfactory/factory"
	. "github.com/takaakit/design-pattern-examples-in-golang/creationalpatterns/abstractfactory/listfactory"
	. "github.com/takaakit/design-pattern-examples-in-golang/creationalpatterns/abstractfactory/tablefactory"

	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"

	"testing"
)

/*
Create a hierarchical link collection as an HTML file. It can be created in either tabular or list format.
*/

func TestMain(m *testing.M) {
	var pb1 *walk.PushButton
	var pb2 *walk.PushButton

	mainWindow := declarative.MainWindow{
		Title:   "Choose 1 or 2",
		MinSize: declarative.Size{Width: 200, Height: 50},
		Size:    declarative.Size{Width: 200, Height: 50},
		Layout:  declarative.VBox{},
		Children: []declarative.Widget{
			declarative.Label{
				Text: "Choose 1 or 2:",
			},
			declarative.PushButton{
				AssignTo: &pb1,
				Text:     "1: Create an HTML file by using ListFactory",
				OnClicked: func() {
					create(NewListFactory())
					os.Exit(0)
				},
			},
			declarative.PushButton{
				AssignTo: &pb2,
				Text:     "2: Create an HTML file by using TableFactory",
				OnClicked: func() {
					create(NewTableFactory())
					os.Exit(0)
				},
			},
		},
	}

	if _, err := mainWindow.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func create(factory Factory) {
	washigtonPost := factory.CreateLink("The Washington Post", "https://www.washingtonpost.com/")
	newYorkTimes := factory.CreateLink("The NewYork Times", "https://www.nytimes.com/")
	financialTimes := factory.CreateLink("The Financial Times", "https://www.ft.com/")

	newspaper := factory.CreateData("Newspaper")
	newspaper.Add(washigtonPost)
	newspaper.Add(newYorkTimes)
	newspaper.Add(financialTimes)

	yahoo := factory.CreateLink("Yahoo!", "https://www.yahoo.com/")
	google := factory.CreateLink("Google", "https://www.google.com/")

	searchEngine := factory.CreateData("Search engine")
	searchEngine.Add(yahoo)
	searchEngine.Add(google)

	linkPage := factory.CreatePage("LinkPage", "James Smith")
	linkPage.Add(newspaper)
	linkPage.Add(searchEngine)

	linkPage.Output(linkPage) // Client-Specified Self pattern
}
