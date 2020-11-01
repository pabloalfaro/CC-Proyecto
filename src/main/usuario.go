package main

//import "fmt"

type Usuario struct{
  username string
  nombre string
  apellidos string
  correo string
  ciudad string
  
}

//Get para username
func GetUsermane(u Usuario) string{
  return u.username
}


//Get y set para nombre
func GetNombreU (u Usuario) string{
  return u.nombre
}

func SetNombreU (u Usuario, n string) {
  u.nombre = n
}


//Get y set para apellidos
func GetApellidosU (u Usuario) string{
  return u.apellidos
}

func SetApellidosU (u Usuario, a string) {
  u.apellidos = a
}


//Get y set para correo
func GetCorreoU (u Usuario) string{
  return u.correo
}

func SetCorreoU (u Usuario, c string) {
  u.correo = c
}


//Get y set para Ciudad
func GetCiudadU (u Usuario) string{
  return u.ciudad
}

func SetCiudadU (u Usuario, c string) {
  u.ciudad = c
}
