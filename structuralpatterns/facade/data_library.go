// ˅
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ˄

type DataLibrary struct {
	// ˅

	// ˄

	// ˅

	// ˄
}

func NewDataLibrary() *DataLibrary {
	// ˅
	if instanceDataLibrary == nil {
		instanceDataLibrary = &DataLibrary{}
	}
	return instanceDataLibrary
	// ˄
}

// Read a data library file.
func (self *DataLibrary) GetProperties(dataLibraryName string) map[string]string {
	// ˅
	file, err := os.Open(dataLibraryName + ".txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var properties = map[string]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "=")
		if len(data) == 2 {
			properties[data[0]] = data[1]
		}
	}

	return properties
	// ˄
}

// ˅
var instanceDataLibrary *DataLibrary = nil

// ˄
