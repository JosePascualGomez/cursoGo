package figuras

import (
	"fmt"
	"math"
)

type Circulo struct {
	Radio float64
}

func (c *Circulo) area() float64 {
	return math.Pow(c.Radio, 2) * math.Pi
}

func (c *Circulo) perimetro() float64 {
	return 2 * math.Pi * c.Radio
}

func (c *Circulo) imprimir() string {
	return fmt.Sprintf("Circulo de radio %f", c.Radio)
}
