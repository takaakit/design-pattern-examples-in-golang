// ˅
package observer

import (
	"fmt"
	"strings"
)

// ˄

// Display values as a bar chart.
type BarChartObserver struct {
	// ˅

	// ˄

	numberSubject *NumberSubject

	// ˅

	// ˄
}

func NewBarChartObserver(numberSubject *NumberSubject) *BarChartObserver {
	// ˅
	return &BarChartObserver{numberSubject: numberSubject}
	// ˄
}

func (b *BarChartObserver) Update(iSubject ISubject) {
	// ˅
	// Before processing, it checks to make sure the changed subject is the subject held.
	if iSubject == &(b.numberSubject.Subject) {
		fmt.Println("Bar chart: " + strings.Repeat("*", b.numberSubject.value))
	}
	// ˄
}

// ˅

// ˄
