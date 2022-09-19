package figuras

import "fmt"

type figura interface {
	area() float64
	perimetro() float64
	imprimir() string
}

func areaFigura(f figura) float64 {
	return f.area()
}

func perimetroFigura(f figura) float64 {
	return f.perimetro()
}

func imprimirFigura(f figura) string {
	return f.imprimir()
}

//Mayuscula para que sea público
func MedidasFigura(f figura) {
	fmt.Println(f.imprimir())
	fmt.Println("El área es:", f.area())
	fmt.Println("El perímetro es:", f.perimetro())
}
