package figuras

import "fmt"

type Cuadrado struct {
	Lado float64
}

// Implementacion
func (c *Cuadrado) area() float64 {
	return 2 * c.Lado
}

func (c *Cuadrado) perimetro() float64 {
	return (4 * c.Lado)
}

func (c *Cuadrado) imprimir() string {
	return fmt.Sprintf("Cuadrado de lado %f", c.Lado)
}
