package main

import (
	"fmt"
	"os"
)

// Create a hierarchical link collection as an HTML file.

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: abstractfactory.exe struct.name.of.ConcreteFactory")
		fmt.Println("Ex.1 : abstractfactory.exe ListFactory")
		fmt.Println("Ex.2 : abstractfactory.exe TableFactory")
	} else {
		factory := getFactory(os.Args[1])

		washigtonPost := factory.CreateLink("The Washington Post", "https://www.washingtonpost.com/")
		newYorkTimes := factory.CreateLink("The NewYork Times", "https://www.nytimes.com/")
		financialTimes := factory.CreateLink("The Financial Times", "https://www.ft.com/")
		yahoo := factory.CreateLink("Yahoo!", "https://www.yahoo.com/")
		google := factory.CreateLink("Google", "https://www.google.com/")
		mlb := factory.CreateLink("MLB", "https://www.mlb.com/")
		fifa := factory.CreateLink("FIFA", "https://www.fifa.com/")
		wba := factory.CreateLink("WBA", "http://www.wbaboxing.com/")
		wbc := factory.CreateLink("WBC", "http://www.wbcboxing.com/")

		newspaper := factory.CreateData("Newspaper")
		newspaper.Add(washigtonPost)
		newspaper.Add(newYorkTimes)
		newspaper.Add(financialTimes)

		searchEngine := factory.CreateData("Search engine")
		searchEngine.Add(yahoo)
		searchEngine.Add(google)

		sports := factory.CreateData("Sports")
		sports.Add(mlb)
		sports.Add(fifa)
		sports.Add(wba)
		sports.Add(wbc)

		linkPage := factory.CreatePage("LinkPage", "James Smith")
		linkPage.Add(newspaper)
		linkPage.Add(searchEngine)
		linkPage.Add(sports)

		context := linkPage.ToHTML()
		linkPage.Output(context)
	}
}

func getFactory(structName string) Factory {
	if structName == "ListFactory" {
		return NewListFactory()
	} else if structName == "TableFactory" {
		return NewTableFactory()
	} else {
		fmt.Println(structName + " is unknown")
		return nil
	}
}
