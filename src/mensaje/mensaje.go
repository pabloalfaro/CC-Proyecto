package mensaje

import "github.com/pabloalfaro/Car-finder/src/usuario"
import "github.com/pabloalfaro/Car-finder/src/anuncio"

type Mensaje struct {
	comprador usuario.Usuario
	anuncio   anuncio.Anuncio
	cuerpo    string
}

//Constructor de mensajes
func NewMensaje(co usuario.Usuario, a anuncio.Anuncio, cu string) Mensaje {
	var mensaje Mensaje

	mensaje.comprador = co
	mensaje.anuncio = a
	mensaje.cuerpo = cu

	return mensaje
}

//Get para vendedor
func GetAnuncioM(m Mensaje) anuncio.Anuncio {
	return m.anuncio
}

//Get para comprador
func GetCompradorM(m Mensaje) usuario.Usuario {
	return m.comprador
}

//Get y set para cuerpo
func GetCuerpoM(m Mensaje) string {
	return m.cuerpo
}
