package main

// Observers observe objects generating a numerical value and display the value.

func main() {
	number := NewRandomNumber()
	digitObserver := NewDigitObserver()
	barChartObserver := NewBarChartObserver()
	number.AddObserver(digitObserver)
	number.AddObserver(barChartObserver)
	number.Generate()
}
