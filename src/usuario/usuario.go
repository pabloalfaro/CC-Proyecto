package usuario

//import "fmt"

type Usuario struct {
	username  string
	nombre    string
	apellidos string
	correo    string
	ciudad    string
}

//Constructor de usuarios
func NewUsuario(u string, n string, a string, co string, ci string) Usuario {
	var usuario Usuario

	usuario.username = u
	usuario.nombre = n
	usuario.apellidos = a
	usuario.correo = co
	usuario.ciudad = ci

	return usuario
}

//Get para username
func GetUsername(u Usuario) string {
	return u.username
}

//Get y set para nombre
func GetNombre(u Usuario) string {
	return u.nombre
}

func SetNombre(u Usuario, n string) {
	u.nombre = n
}

//Get y set para apellidos
func GetApellidos(u Usuario) string {
	return u.apellidos
}

func SetApellidos(u Usuario, a string) {
	u.apellidos = a
}

//Get y set para correo
func GetCorreo(u Usuario) string {
	return u.correo
}

func SetCorreo(u Usuario, c string) {
	u.correo = c
}

//Get y set para Ciudad
func GetCiudad(u Usuario) string {
	return u.ciudad
}

func SetCiudadU(u Usuario, c string) {
	u.ciudad = c
}
