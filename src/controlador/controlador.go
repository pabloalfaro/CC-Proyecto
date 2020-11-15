package controlador

//import "fmt"

import "github.com/pabloalfaro/Car-finder/src/anuncio"
import "github.com/pabloalfaro/Car-finder/src/coche"
//import "github.com/pabloalfaro/Car-finder/src/mensaje"
import "github.com/pabloalfaro/Car-finder/src/usuario"

type Controlador struct{
//El  identificador será el propio Username
	usuarios map[string]usuario.Usuario

	coches []coche.Coche
	anuncios []anuncio.Anuncio
}


func NuevoControlador() Controlador{
	//Inicializo el map
	var cont Controlador
	cont.usuarios = make(map[string]usuario.Usuario)
	
	return cont
}

//Devuelve un boolean igual a True si se ha hecho la inserción y False si no se ha podido hacer
func (cont *Controlador) NuevoUsuario(u string, n string, a string, co string, ci string) bool {
	//Compruebo si ya existe un usuario con ese username
	_, ok := cont.usuarios[u]

	//Si la variable ok es igual a true entonces el username ya está registrado
	if ok {
		return false
	} else {
		var user usuario.Usuario
		user = usuario.NewUsuario(u, n, a, co, ci)
		cont.usuarios[u] = user

		//Se ha podido hacer la inserción
		return true
	}
}

func (cont *Controlador) NuevoAnuncio(u string, p float32, coc coche.Coche, k int, e string, ciu string, d string, col string) bool{
	_, ok := cont.usuarios[u]
	
	if ok {
		var anun anuncio.Anuncio

		anun = anuncio.NewAnuncio(u, p, coc, k, e, ciu, d, col)
		cont.anuncios = append(cont.anuncios, anun)
		return true
	} else {
		return false
	}
}
