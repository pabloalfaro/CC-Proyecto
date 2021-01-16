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



//Devuelve un boolean igual a True si se ha encontrado y False si no
func (cont *Controlador) BuscarUsuario(u string) (usuario.Usuario, bool) {
	//Compruebo si ya existe un usuario con ese username
	user, ok := cont.usuarios[u]

	//Si la variable ok es igual a true entonces el username ya está registrado
	if ok {
		return user, true
	} else {
		return user, false
	}
}



func (cont *Controlador) ModificarUsuario(u string, n string, a string, co string, ci string) (bool) {
	//Compruebo si ya existe un usuario con ese username
	user, ok := cont.usuarios[u]

	//Si la variable ok es igual a true entonces el username ya está registrado
	if ok {
		if n != ""{
			usuario.SetNombre(user, n)
		}
		if a != ""{
			usuario.SetNombre(user, a)
		}
		if co != ""{
			usuario.SetNombre(user, co)
		}
		if ci != ""{
			usuario.SetNombre(user, ci)
		}
		return true
	} else {
		return false
	}
}



func (cont *Controlador) BorrarUsuario(u string) (bool) {
	//Compruebo si ya existe un usuario con ese username
	_, ok := cont.usuarios[u]
	
	if ok {
		delete(cont.usuarios, u)
		return true
	}else{
		return false
	}
}



func (cont *Controlador) NuevoCoche(ma string, mo string, se string, po int) bool{
	_, r := cont.BuscarCoche(ma, mo, se, po)
	
	if !r {
		var car coche.Coche

		car = coche.NewCoche(ma, mo, se, po)
		cont.coches = append(cont.coches, car)
		return true
	}else{
		return false
	}
}



func (cont *Controlador) BuscarCoche(ma string, mo string, se string, po int) (int, bool){
	var car coche.Coche

	car = coche.NewCoche(ma, mo, se, po)
	
	for index, i := range cont.coches {
		if i == car{
			return index, true
		}
	}
	return -1, false
}



func (cont *Controlador) BorrarCoche(ma string, mo string, se string, po int) bool{
	index, r := cont.BuscarCoche(ma, mo, se, po)
	
	if r {
		cont.coches = append(cont.coches[index:], cont.coches[index+1:]...)
		return true
	}else{
		return false
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
