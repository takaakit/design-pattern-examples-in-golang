package main

// Someone handles a trouble.

func main() {
	emily := NewLazySupporter("Emily")
	william := NewMoodySupporter("William")
	amelia := NewSpecialSupporter("Amelia", 153)
	michael := NewSpecialSupporter("Michael", 340)
	joseph := NewLimitedSupporter("Joseph", 250)
	lily := NewLimitedSupporter("Lily", 350)

	// Make a chain.
	emily.SetNext(amelia).SetNext(michael).SetNext(joseph).SetNext(lily).SetNext(william)

	// Various troubles occurred.
	for i := 0; i < 500; i += 17 {
		emily.Support(emily, NewTrouble(i))
	}
}
