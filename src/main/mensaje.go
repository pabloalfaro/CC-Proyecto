package main

//import "fmt"

type Mensaje struct{
  comprador Usuario
  cuerpo string
  
}

//Get y set para nombre
func GetCompradorM (m Mensaje) Usuario{
  return m.comprador
}

func SetCompradorM (m Mensaje, u Usuario) {
  m.comprador = u
}


//Get y set para cuerpo
func GetCuerpoM (m Mensaje) string{
  return m.cuerpo
}

func SetCuerpoM (m Mensaje, c string) {
  m.cuerpo = c
}
