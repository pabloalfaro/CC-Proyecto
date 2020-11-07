package mensaje

//import "fmt"

type Mensaje struct{
  comprador Usuario
  anuncio Anuncio
  cuerpo string
  
}


//Constructor de mensajes
func NewMensaje (co Usuario, a Anuncio, cu string) Mensaje{
  var mensaje Mensaje
  
  mensaje.comprador = co
  mensaje.anuncio = a
  mensaje.cuerpo = cu
  
  return mensaje
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


