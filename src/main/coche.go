package main

//import "fmt"

type Coche struct{
  marca string
  modelo string
  serie string
  potencia int

}

//Get y set para marca
func GetMarcaC (c Coche) string{
  return c.marca
}

func SetMarcaC (c Coche, m string) {
  c.marca = m
}


//Get y set para modelo
func GetModeloC (c Coche) string{
  return c.modelo
}

func SetModeloC (c Coche, m string) {
  c.modelo = m
}


//Get y set para serie
func GetSerieC (c Coche) string{
  return c.serie
}

func SetSerieC (c Coche, s string) {
  c.serie = s
}


//Get y set para potencia
func GetPotenciaC (c Coche) int{
  return c.potencia
}

func SetPotenciaC (c Coche, p int) {
  c.potencia = p
}

