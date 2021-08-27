[<img src="./screenshots/DiagramMap.svg">](https://raw.githubusercontent.com/takaakit/design-pattern-examples-in-golang/master/screenshots/DiagramMap.svg)

Design Pattern Examples in Golang
===

Model and code examples of GoF Design Patterns for Golang.  
This project is available for the following objectives:  

* To understand GoF Design Pattern examples in Golang.
* To understand the mapping between UML model and Golang code.
* To try Model-Driven Development (MDD) using Astah and M PLUS plug-in.

> UML model example:

![](screenshots/CompositePattern.svg "Composite Pattern")

<a id="code-example"></a>
> Golang code example:

```golang:File class
// ˅
package bridge

import (
	"fmt"
	"strconv"
)

// ˄

type File struct {
	// ˅

	// ˄

	name string

	size int

	// ˅

	// ˄
}

func NewFile(name string, size int) *File {
	// ˅
	return &File{name: name, size: size}
	// ˄
}

func (f *File) GetName() string {
	// ˅
	return f.name
	// ˄
}

func (f *File) GetSize() int {
	// ˅
	return f.size
	// ˄
}

// Print this element with the "upperPath".
func (f *File) Print(upperPath string) {
	// ˅
	fmt.Println(upperPath + "/" + f.String())
	// ˄
}

func (f *File) String() string {
	// ˅
	return f.GetName() + " (" + strconv.Itoa(f.GetSize()) + ")"
	// ˄
}

// ˅

// ˄
```

Installation on Windows
------------
This project uses the [Walk library](https://github.com/lxn/walk), which works only on Windows.

**Astah**
* Download Astah UML or Professional from [Change Vision site](http://astah.net/download), and install.  

**M PLUS plug-in**
* Download [M PLUS plug-in](https://sites.google.com/view/m-plus-plugin/download) and install. Use **ver.2.7.6** or higher.

**Development Environment for Golang**
* Install [Golang binary](https://golang.org/dl/).
* Run `go get github.com/go-delve/delve/cmd/dlv` to install [delve](https://github.com/derekparker/delve/blob/master/Documentation/installation/README.md).
* Run `go get github.com/lxn/walk` to install [Walk library](https://github.com/lxn/walk).
* Install [VS Code](https://code.visualstudio.com/download), add [Golang extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go).
* Launch VS Code. Click **View -> Command Pallete (Cmd+Shift+P)** and run `Go: Install/Update Tools`. Check all dependencies and click OK.

Usage on Windows
-----
This project uses the [Walk library](https://github.com/lxn/walk), which works only on Windows.

**Code Generation from UML**
  1. Open the Astah file (model/DesignPatternExamplesInGolang.asta).
  2. Select model elements on the model browser of Astah.
  3. Click the **Generate Code** button.  
  ![](screenshots/GenerateCode.gif "Generate Code")  
  The generated code has **User Code Area**. The User Code Area is the area enclosed by "˅" and "˄". Handwritten code written in the User Code Area remains after a re-generation. [View code example](#code-example).  
  For detailed usage of the tools, please see [Astah Manual](http://astah.net/manual) and [M PLUS plug-in Tips](https://sites.google.com/view/m-plus-plugin-tips).

**Build and Run (in Debug Mode)**
  1. Open the workspace file (design-pattern-examples-in-golang.code-workspace) in VS Code.
  2. Select the **Run and Debug** icon in the Activity Bar on the side of VS Code.
  3. Select the configuration named Debug Program using the Configuration drop-down in the Run view.
  4. Click the **Start Debugging** button.  
     ![](screenshots/BuildAndRun.gif "Build and Run")  

References
----------
* Gamma, E. et al. Design Patterns: Elements of Reusable Object-Oriented Software, Addison-Wesley, 1994
* Hiroshi Yuki. Learning Design Patterns in Java [In Japanese Language], Softbank publishing, 2004
* Schmager, F. Evaluating the GO Programming Language with Design Patterns, 2010

Licence
-------
This project is licensed under the Creative Commons Zero (CC0) license. The model and code are completely free to use.

[![CC0](http://i.creativecommons.org/p/zero/1.0/88x31.png "CC0")](http://creativecommons.org/publicdomain/zero/1.0/deed)
