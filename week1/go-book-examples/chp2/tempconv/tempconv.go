package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

func (c Celsius) String() string    { return fmt.Sprintf("%gC", float64(c)) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%gF", float64(f)) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", float64(k)) }
