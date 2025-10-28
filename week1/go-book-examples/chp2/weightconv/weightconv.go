package weightconv

import "fmt"

type Pound float64
type Kilogram float64

func (p Pound) String() string    { return fmt.Sprintf("%gpound", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }

func PToKg(p Pound) Kilogram { return Kilogram(p / 0.453) }
func KgToP(k Kilogram) Pound { return Pound(k * 2.205) }
