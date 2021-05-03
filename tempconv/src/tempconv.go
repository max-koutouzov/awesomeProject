package src

import (
	"flag"
	"fmt"
)

type Celsius float64
type Farenheit float64

func CToF(c Celsius) Farenheit {
	return Farenheit(c*9.0/5.0 + 32.0)
}

func FToC(f Farenheit) Celsius {
	return Celsius(f-32.0) * 5.0 / 9.0
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

// *celsiusFlag satisfies the flag.Value interface.

type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Scanf(s, "%f%s", &value, &unit)

	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Farenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperatures %q", s)
}

//!-celsiusFlag

//!+CelsiusFlag

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C".

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

//!-CelsiusFlag
