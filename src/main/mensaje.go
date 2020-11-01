package main

//import "fmt"

type Mensaje struct{
  comprador Usuario
  anuncio Anuncio
  cuerpo string
  
}

//Get para vendedor
func GetAnuncioM (m Mensaje) Anuncio{
  return m.anuncio
}


//Get para comprador
func GetCompradorM (m Mensaje) Usuario{
  return m.comprador
}


//Get y set para cuerpo
func GetCuerpoM (m Mensaje) string{
  return m.cuerpo
}


