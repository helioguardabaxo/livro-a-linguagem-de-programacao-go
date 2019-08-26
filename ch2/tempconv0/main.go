package tempconv0

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

// CToF Converts Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC Converts Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
