// ˅
package main

import (
	"fmt"
	"strconv"
	"time"
)

// ˄

// Display values with digits.
type DigitObserver struct {
	// ˅

	// ˄

	// ˅

	// ˄
}

func NewDigitObserver() *DigitObserver {
	// ˅
	return &DigitObserver{}
	// ˄
}

func (self *DigitObserver) Update(number *Number) {
	// ˅
	fmt.Println("Digit    : " + strconv.Itoa(number.value))
	time.Sleep(100 * time.Millisecond)
	// ˄
}

// ˅

// ˄
