package lengthconv

import "fmt"

type Foot float64
type Meter float64

func (f Foot) String() string  { return fmt.Sprintf("%gFeet", float64(f)) }
func (m Meter) String() string { return fmt.Sprintf("%gMeter", float64(m)) }

func FToM(f Foot) Meter { return Meter(f / 0.3048) }
func MToF(m Meter) Foot { return Foot(m * 0.3048) }
