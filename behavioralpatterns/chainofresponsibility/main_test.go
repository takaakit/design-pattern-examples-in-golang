package chainofresponsibility

import "testing"

/*
A trouble is turned around among supporters, and the trouble will be handled by the supporter who can handle it.
There are four types of supporters below:
* LazySupporter doesn't handle any trouble
* MoodySupporter handles odd ID troubles
* SpecialSupporter handles a trouble of the target ID.
* LimitedSupporter handles troubles below the limit ID.
*/

func Test(t *testing.T) {
	emily := NewLazySupporter("Emily")
	william := NewMoodySupporter("William")
	amelia := NewSpecialSupporter("Amelia", 6)
	joseph := NewLimitedSupporter("Joseph", 5)

	// Make a chain.
	emily.Next = william
	william.Next = amelia
	amelia.Next = joseph

	// Various troubles occurred.
	for i := 0; i < 10; i++ {
		emily.Support(emily, NewTrouble(i))
	}
}
