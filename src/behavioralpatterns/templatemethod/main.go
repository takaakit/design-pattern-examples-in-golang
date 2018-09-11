package main

// Display the character and string repeatedly 5 times.

func main() {
	display1 := NewCharDisplay("H")                   // Create an instance of the CharDisplay
	display2 := NewStringDisplay("Hello world.")      // Create an instance of the StringDisplay
	display3 := NewStringDisplay("Nice to meet you.") // Create an instance of the StringDisplay

	// Any instance can be called with the same method name
	display1.Output(display1)
	display2.Output(display2)
	display3.Output(display3)
}
