package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Analyze the syntax of the mini-language composed of "forward", "left", "right", and "repeat" commands.

-----
Examples before and after syntax analysis.
* Ex.1
Before : "program end"
After  : [program []]

* Ex.2
Before : "program forward right left end"
After  : [program [forward, right, left]]

* Ex.3
Before : "program repeat 4 forward right end end"
After  : [program [[repeat 4 [forward, right]]]]
*/

func main() {
	var fileName = "program.txt"
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(fileName + " could not read")
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var line = scanner.Text()
		fmt.Println("TEXT > \"" + line + "\"")
		node := NewProgram()
		node.Parse(NewContext(line))
		fmt.Println("NODE > " + node.ToString())
	}
	if serr := scanner.Err(); serr != nil {
		fmt.Println(fileName + " scan error")
	}
}
