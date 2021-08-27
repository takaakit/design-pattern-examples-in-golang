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

func TestMain(m *testing.M) {
	p := NewPrintMessageDisplay("Nice to meet you")
	p.PrintWeak()
	p.PrintStrong()
}
