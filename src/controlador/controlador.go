package controlador

import "fmt"

import "github.com/pabloalfaro/Car-finder/src/anuncio"
import "github.com/pabloalfaro/Car-finder/src/coche"
//import "github.com/pabloalfaro/Car-finder/src/mensaje"
import "github.com/pabloalfaro/Car-finder/src/usuario"


//El  identificador será el propio Username
var Usuarios map [string] usuario.Usuario

var Coches [] coche.Coche

var Anuncios [] anuncio.Anuncio

func Main(){
	//Inicializo el map
	Usuarios = make (map [string] usuario.Usuario)
	
	fmt.Println("Usuarios iniciado")
}

//Devuelve un boolean igual a True si se ha hecho la inserción y False si no se ha podido hacer
func NuevoUsuario (u string, n string, a string, co string, ci string) bool{
  //Compruebo si ya existe un usuario con ese username
  _, ok := Usuarios[u]
  
  //Si la variable ok es igual a true entonces el username ya está registrado
  if (ok){
    return false
  }else{
    var user usuario.Usuario
    user = usuario.NewUsuario(u, n, a, co, ci)
    Usuarios [u] = user
    
    //Se ha podido hacer la inserción
    return true
  }
}


func NuevoAnuncio (u usuario.Usuario, p float32, coc coche.Coche, k int, e string, ciu string, d string, col string) {
  var anun anuncio.Anuncio
  
	anun = anuncio.NewAnuncio(u, p, coc, k, e, ciu, d, col)
	Anuncios = append(Anuncios, anun)
}
