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

func TestMain(m *testing.M) {
	numberSubject := NewNumberSubject()
	numberSubject.Attach(NewDigitObserver(numberSubject))
	numberSubject.Attach(NewBarChartObserver(numberSubject))

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		numberSubject.SetValue(rand.Intn(50))
		time.Sleep(200 * time.Millisecond)
	}
}
