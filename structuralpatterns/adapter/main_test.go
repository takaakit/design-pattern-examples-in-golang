package main

import (
	"testing"
)

/*
Display the given string as follows
```
-- Nice to meet you --
```
or display it as follows.
```
[[ Nice to meet you ]]
```
*/

func Test(t *testing.T) {
	p := NewPrintMessageDisplay("Nice to meet you")
	p.PrintWeak()
	p.PrintStrong()
}
