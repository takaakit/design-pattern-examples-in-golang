package main

// Display a character string with a decorative frame.

func main() {
	displayA := NewMessageDisplay("Nice to meet you.")
	displayA.Show(displayA)

	displayB := NewSideFrame(displayA, "!")
	displayB.Show(displayB)

	displayC := NewFullFrame(displayB)
	displayC.Show(displayC)

	displayD := NewSideFrame(
		NewFullFrame(
			NewFullFrame(
				NewSideFrame(
					NewSideFrame(
						NewFullFrame(
							NewMessageDisplay("See you again.")),
						"#"),
					"#"))),
		"#")
	displayD.Show(displayD)
}
