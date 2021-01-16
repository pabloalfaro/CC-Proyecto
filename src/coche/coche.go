package coche

//import "fmt"

type Coche struct {
	marca    string
	modelo   string
	serie    string
	potencia int
}

func NewCoche(ma string, mo string, s string, p int) Coche {
	var coche Coche

	coche.marca = ma
	coche.modelo = mo
	coche.serie = s
	coche.potencia = p

	return coche
}

//Get para marca
func GetMarca(c Coche) string {
	return c.marca
}

//Get para modelo
func GetModelo(c Coche) string {
	return c.modelo
}

//Get para serie
func GetSerie(c Coche) string {
	return c.serie
}

//Get para potencia
func GetPotencia(c Coche) int {
	return c.potencia
}
