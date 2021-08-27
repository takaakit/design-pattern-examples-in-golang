package interpreter

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

/*
An interpreter for mini language to operate radio controlled car. It parses the following syntax
composed of "forward", "left", "right", and "repeat" commands:
```
<program>      ::= program <command list>
<command list> ::= <command>* end
<command>      ::= <repeat> | <action>
<repeat>       ::= repeat <number> <command list>
<action>       ::= forward | right | left
<number>       ::= 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9
```
___
Examples before and after syntax analysis.

Ex.1
```
Before parsing : program end
After parsing  : [program []]
```

Ex.2
```
Before parsing : program forward right left end
After parsing  : [program [forward, right, left]]
```

Ex.3
```
Before parsing : program repeat 4 forward right end end
After parsing  : [program [repeat 4 [forward, right]]]
```
*/

func TestMain(m *testing.M) {
	var fileName = "program.txt"
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(fileName + " could not read")
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println("Before parsing : " + text)
		node := NewProgram()
		node.Parse(NewContext(text))
		fmt.Println("After parsing  : " + node.String())
	}
	if serr := scanner.Err(); serr != nil {
		fmt.Println("Failed to read file: " + fileName)
	}
}
