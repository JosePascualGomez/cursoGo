package figuras

import "fmt"

type Rectangulo struct {
	Ancho float64
	Alto  float64
}

func (c *Rectangulo) area() float64 {
	return c.Ancho * c.Alto
}

func (c *Rectangulo) perimetro() float64 {
	return (2 * c.Ancho) + (2 * c.Alto)
}

func (c *Rectangulo) imprimir() string {
	return fmt.Sprintf("Rectangulo de alto %f y ancho %f", c.Alto, c.Ancho)
}
