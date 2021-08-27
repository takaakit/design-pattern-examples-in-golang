// ˅
package observer

import (
	"fmt"
	"strconv"
)

// ˄

// Display values as a number.
type DigitObserver struct {
	// ˅

	// ˄

	numberSubject *NumberSubject

	// ˅

	// ˄
}

func NewDigitObserver(numberSubject *NumberSubject) *DigitObserver {
	// ˅
	return &DigitObserver{numberSubject: numberSubject}
	// ˄
}

func (d *DigitObserver) Update(iSubject ISubject) {
	// ˅
	if iSubject == &(d.numberSubject.Subject) {
		fmt.Println("Digit    : " + strconv.Itoa(d.numberSubject.value))
	}
	// ˄
}

// ˅

// ˄
