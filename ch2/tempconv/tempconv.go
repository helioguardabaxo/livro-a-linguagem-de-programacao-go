package tempconv

import "fmt"

// Celsius ...
type Celsius float64

// Fahrenheit ...
type Fahrenheit float64

const (
	// AbsoluteZeroC ...
	AbsoluteZeroC Celsius = -273.15
	// FreezingC ...
	FreezingC Celsius = 0
	// BoilingC ...
	BoilingC Celsius = 100
)

// Converts Celsius to Fahrenheit
func (c Celsius) String() string { return fmt.Sprintf("%gºC", c) }

// Converts Fahrenheit to Celsius
func (f Fahrenheit) String() string { return fmt.Sprintf("%gºF", f) }
