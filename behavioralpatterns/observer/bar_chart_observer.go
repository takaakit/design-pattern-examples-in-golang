// ˅
package main

import (
	"fmt"
	"time"
)

// ˄

// Display values with a bar chart.
type BarChartObserver struct {
	// ˅

	// ˄

	// ˅

	// ˄
}

func NewBarChartObserver() *BarChartObserver {
	// ˅
	return &BarChartObserver{}
	// ˄
}

func (self *BarChartObserver) Update(number *Number) {
	// ˅
	fmt.Print("Bar chart: ")
	for i := 0; i < number.value; i++ {
		fmt.Print("*")
	}
	fmt.Println("")
	time.Sleep(100 * time.Millisecond)
	// ˄
}

// ˅

// ˄
