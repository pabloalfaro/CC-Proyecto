package anuncio

import "github.com/pabloalfaro/Car-finder/src/coche"

type Anuncio struct{
  precio float32
  coche coche.Coche
  km int
  estado string
  ciudad string
  descripcion string
  color string

}


//Constructor anuncio
func NewAnuncio(p float32, coc coche.Coche, k int, e string, ciu string, d string, col string) Anuncio{
  var anuncio Anuncio
  
  anuncio.precio = p
  anuncio.coche = coc
  anuncio.km = k
  anuncio.estado = e
  anuncio.ciudad = ciu
  anuncio.descripcion = d
  anuncio.color = col
  
  return anuncio
}

//Get y set para Precio
func GetPrecioA (a Anuncio) float32{
  return a.precio
}

func SetPrecioA (a Anuncio, p float32) {
  a.precio = p
}


//Get para Coche
func GetCocheA (a Anuncio) coche.Coche{
  return a.coche
}

//Get y set para km
func GetKmA (a Anuncio) int{
  return a.km
}

func SetKmA (a Anuncio, k int) {
  a.km = k
}


//Get y set para estado
func GetEstadoA (a Anuncio) string{
  return a.estado
}

func SetEstadoA (a Anuncio, e string) {
  a.estado = e
}


//Get y set para ciudad
func GetCiudadA (a Anuncio) string{
  return a.ciudad
}

func SetCiudadA (a Anuncio, c string) {
  a.ciudad = c
}


//Get y set para descripcion
func GetDescripcionA (a Anuncio) string{
  return a.descripcion
}

func SetDescripcionA (a Anuncio, d string) {
  a.descripcion = d
}


//Get y set para color
func GetColorA (a Anuncio) string{
  return a.color
}

func SetColorA (a Anuncio, c string) {
  a.color = c
}

