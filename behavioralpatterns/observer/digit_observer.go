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
	// Before processing, it checks to make sure the changed subject is the subject held.
	if iSubject == &(d.numberSubject.Subject) {
		fmt.Println("Digit    : " + strconv.Itoa(d.numberSubject.value))
	}
	// ˄
}

// ˅

// ˄
