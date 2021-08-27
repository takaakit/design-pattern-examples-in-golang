package bridge

import (
	"fmt"
	"testing"
)

/*
Represents a file system composed of files and directories. FileSystemElement makes
it possible to treat File and Directory uniformly.
*/

func TestMain(m *testing.M) {
	fmt.Println("Create a file system...")

	binDir := NewDirectory("bin")
	lsFile := NewFile("ls", 20)
	binDir.Add(lsFile)
	mkdirFile := NewFile("mkdir", 40)
	binDir.Add(mkdirFile)

	emilyDir := NewDirectory("emily")
	homeworkFile := NewFile("homework.doc", 60)
	emilyDir.Add(homeworkFile)

	jamesDir := NewDirectory("james")
	appFile := NewFile("app.exe", 80)
	jamesDir.Add(appFile)

	homeDir := NewDirectory("home")
	homeDir.Add(emilyDir)
	homeDir.Add(jamesDir)

	rootDir := NewDirectory("root")
	rootDir.Add(homeDir)
	rootDir.Add(binDir)

	rootDir.Print("")
}
