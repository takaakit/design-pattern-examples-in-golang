package observer

import (
	"math/rand"
	"testing"
	"time"
)

/*
Observers observe a Subject object holding a numerical value and display the value.
The display formats are digits and bar charts.
*/

func Test(t *testing.T) {
	numberSubject := NewNumberSubject()
	numberSubject.Attach(NewDigitObserver(numberSubject))
	numberSubject.Attach(NewBarChartObserver(numberSubject))

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		numberSubject.SetValue(rand.Intn(50))
		time.Sleep(200 * time.Millisecond)
	}
}
