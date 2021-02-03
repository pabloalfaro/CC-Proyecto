package controlador

import (
	"github.com/pabloalfaro/Car-finder/src/anuncio"
	"github.com/pabloalfaro/Car-finder/src/coche"
	"github.com/pabloalfaro/Car-finder/src/usuario"
	"database/sql"
  _ "github.com/go-sql-driver/mysql"
  "fmt"
  "log"
)

//Conexión con la base de datos
var DbUser = "docker"
var DbHost = "godockerDB"
var DbPassword = "password"
var DbName = "carfinder"
var DbDriver = "mysql"
var URL = fmt.Sprintf("%s:%s@tcp(%s)/%s", DbUser, DbPassword, DbHost, DbName)
var db, err = sql.Open(DbDriver, URL)


type Controlador struct{
//El  identificador será el propio Username
	usuarios map[string]usuario.Usuario
	coches map[int]coche.Coche
	anuncios map[int]anuncio.Anuncio
	
}


func NuevoControlador() Controlador{
	//Inicializo el map
	var cont Controlador
	cont.usuarios = make(map[string]usuario.Usuario)
	cont.coches = make(map[int]coche.Coche)
	cont.anuncios = make(map[int]anuncio.Anuncio)
	
	if err != nil {
		log.Print(err.Error())
	}
	
	return cont
}

//Devuelve un boolean igual a True si se ha hecho la inserción y False si no se ha podido hacer
func (cont *Controlador) NuevoUsuario(u string, n string, a string, co string, ci string) bool {
	consulta := fmt.Sprintf("INSERT INTO usuarios (username, nombre, apellido, correo, ciudad) VALUES ('%s', '%s', '%s', '%s', '%s');", u, n, a, co, ci)
	insert, err := db.Query(consulta)
	if err != nil {
		log.Print(err.Error())
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
	defer insert.Close()
	var user usuario.Usuario
	user = usuario.NewUsuario(u, n, a, co, ci)
	cont.usuarios[u] = user

	//Se ha podido hacer la inserción
	return true
}



//Devuelve un boolean igual a True si se ha encontrado y False si no
func (cont *Controlador) BuscarUsuario(u string) (usuario.Usuario, bool) {
	//Compruebo si ya existe un usuario con ese username
	//user, _ := cont.usuarios[u]
	
	var user usuario.Usuario
	consulta := fmt.Sprintf("SELECT * FROM usuarios WHERE username = '%s'", u)
	row:= db.QueryRow(consulta);
	var n, a, co, ci string
	err = row.Scan(&u, &n, &a, &co, &ci) 
	if err != nil{
			log.Print(err.Error())
			user, ok := cont.usuarios[u]
			//Si la variable ok es igual a true entonces el username ya está registrado
			if ok {
				return user, true
			} else {
				return user, true
			}
		}else{
		user = usuario.NewUsuario(u, n, a, co, ci)
		return user, true
	}
}



func (cont *Controlador) ModificarUsuario(u string, n string, a string, co string, ci string) (bool) {
	user, sol := cont.BuscarUsuario(u)
	if !sol {
		return false
	}else{
		if n == ""{
			n = usuario.GetNombre(user)
		}
		if a == ""{
			a = usuario.GetApellidos(user)
		}
		if co == ""{
			co = usuario.GetCorreo(user)
		}
		if ci == ""{
			ci = usuario.GetCiudad(user)
		}
		consulta := fmt.Sprintf("UPDATE usuarios SET nombre='%s', apellido='%s', correo='%s', ciudad='%s' WHERE username='%s'",n,a,co,ci,u)
		_, err := db.Exec(consulta)
		if err!=nil{
			log.Print(err.Error())
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
		}
		return true
	}
}



func (cont *Controlador) BorrarUsuario(u string) (bool) {
	consulta := fmt.Sprintf("DELETE FROM usuarios WHERE username='%s'", u)
	_, err := db.Exec(consulta)
	if err!=nil{
		log.Print(err.Error())
		_, ok := cont.usuarios[u]
	
		if ok {
			delete(cont.usuarios, u)
			return true
		}else{
			return false
		}
	}
	return true
}



func (cont *Controlador) NuevoCoche(ma string, mo string, se string, po int) bool{
	consulta := fmt.Sprintf("INSERT INTO coches (marca, modelo, serie, potencia) VALUES ('%s', '%s', '%s', '%d');", ma, mo, se, po)
	insert, err := db.Query(consulta)
	if err != nil {
		log.Print(err.Error())
		_, r := cont.BuscarCoche(1)
	
		if !r {
			var car coche.Coche

			car = coche.NewCoche(1, ma, mo, se, po)
			cont.coches[1] =  car
			return true
		}else{
			return false
		}
	}
	defer insert.Close()

	//Se ha podido hacer la inserción
	return true
}



func (cont *Controlador) BuscarCoche(id  int) (coche.Coche, bool){
	var car coche.Coche
	consulta := fmt.Sprintf("SELECT * FROM coches WHERE id = '%d'", id)
	row:= db.QueryRow(consulta);
	var ma, mo, se string
	var po int
	err = row.Scan(&id, &ma, &mo, &se, &po) 
	if err != nil{
			log.Print(err.Error())
			
			car, ok := cont.coches[id]
			if ok {
				return car, true
			} else {
				return car, false
			}
		}else{
		car = coche.NewCoche(id, ma, mo, se, po)
		return car, true
	}
}



func (cont *Controlador) BorrarCoche(id int) bool{
	consulta := fmt.Sprintf("DELETE FROM coches WHERE id='%d'", id)
	_, err := db.Exec(consulta)
	if err!=nil{
		log.Print(err.Error())
		_, ok := cont.coches[id]
	
		if ok {
			delete(cont.coches, id)
			return true
		}else{
			return false
		}
	}
	return true
}



func (cont *Controlador) NuevoAnuncio(u string, p float32, coc int, k int, e string, ciu string, d string, col string) bool{
	consulta := fmt.Sprintf("INSERT INTO anuncios (usuario, precio, coche, km, estado, ciudad, descripcion, color) VALUES ('%s', '%g', '%d', '%d', '%s', '%s', '%s', '%s');", u, p, coc, k, e, ciu, d, col)
	insert, err := db.Query(consulta)
	if err != nil {
		log.Print(err.Error())
		_, ok := cont.usuarios[u]
		if ok {
			var anun anuncio.Anuncio
			anun = anuncio.NewAnuncio(1, u, p, coc, k, e, ciu, d, col)
			cont.anuncios[1] = anun
			return true
		} else {
			return false
		}
	}
	defer insert.Close()

	//Se ha podido hacer la inserción
	return true
}



func (cont *Controlador) BuscarAnuncio(id int) (anuncio.Anuncio, bool){
	var anun anuncio.Anuncio
	consulta := fmt.Sprintf("SELECT * FROM anuncios WHERE id = '%d';", id)
	row:= db.QueryRow(consulta);
	var u, e, ciu, d, col string
	var coc, k int
	var p float32  
	err = row.Scan(&id, &u, &p, &coc, &k, &e, &ciu, &d, &col) 
	if err != nil{
			log.Print(err.Error())
			anun, ok := cont.anuncios[id]
			if ok {
				return anun, true
			} else {
				return anun, false
			}
			
		}else{
		anun = anuncio.NewAnuncio(id, u, p, coc, k, e, ciu, d, col)
		return anun, true
	}
}



func (cont *Controlador) BorrarAnuncio(id int) bool{
	consulta := fmt.Sprintf("DELETE FROM anuncios WHERE id='%d'", id)
	_, err := db.Exec(consulta)
	if err!=nil{
		log.Print(err.Error())
		_, ok := cont.anuncios[id]
	
		if ok {
			delete(cont.anuncios, id)
			return true
		}else{
			return false
		}
	}
	return true
}
