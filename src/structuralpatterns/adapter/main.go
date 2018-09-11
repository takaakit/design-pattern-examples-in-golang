package main

/*
Display a character string as follows
-- Nice to meet you --
or display it as follows.
[[ Nice to meet you ]]
*/

func main() {
	p := NewPrintMessageDisplay("Nice to meet you")
	p.PrintWeak()
	p.PrintStrong()
}
