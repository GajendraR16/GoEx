package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/GajendraR16/chp2/lengthconv"
	"github.com/GajendraR16/chp2/tempconv"
	"github.com/GajendraR16/chp2/weightconv"
)

var (
	celsius = flag.Float64("c", 0.0, "temperature in Celsius")
	fahr    = flag.Float64("f", 0.0, "temperature in Fahrenheit")
	kelvin  = flag.Float64("k", 0.0, "temperature in Kelvin")

	pound    = flag.Float64("p", 0.0, "weight in Pounds")
	kilogram = flag.Float64("kg", 0.0, "weight in Kilograms")

	meter = flag.Float64("m", 0.0, "length in Meters")
	foot  = flag.Float64("ft", 0.0, "length in Feet")
)

func main() {
	flag.Parse()

	switch {
	// Temperature conversions
	case *celsius != 0.0:
		c := tempconv.Celsius(*celsius)
		fmt.Printf("%.2f°C = %.2f°F = %.2fK\n", c, tempconv.CToF(c), tempconv.CToK(c))
	case *fahr != 0.0:
		f := tempconv.Fahrenheit(*fahr)
		fmt.Printf("%.2f°F = %.2f°C = %.2fK\n", f, tempconv.FToC(f), tempconv.FToK(f))
	case *kelvin != 0.0:
		k := tempconv.Kelvin(*kelvin)
		fmt.Printf("%.2fK = %.2f°C = %.2f°F\n", k, tempconv.KToC(k), tempconv.KToF(k))

	// Weight conversions
	case *pound != 0.0:
		p := weightconv.Pound(*pound)
		fmt.Printf("%.2f lb = %.2f kg\n", p, weightconv.PToKg(p))
	case *kilogram != 0.0:
		kg := weightconv.Kilogram(*kilogram)
		fmt.Printf("%.2f kg = %.2f lb\n", kg, weightconv.KgToP(kg))

	// Length conversions
	case *meter != 0.0:
		m := lengthconv.Meter(*meter)
		fmt.Printf("%.2f m = %.2f ft\n", m, lengthconv.MToF(m))
	case *foot != 0.0:
		ft := lengthconv.Foot(*foot)
		fmt.Printf("%.2f ft = %.2f m\n", ft, lengthconv.FToM(ft))

	default:
		fmt.Fprintln(os.Stderr, "Usage: go run uc.go -c 100 | -f 212 | -k 373 | -p 100 | -kg 45 | -m 10 | -ft 3.28")
		os.Exit(1)
	}
}
