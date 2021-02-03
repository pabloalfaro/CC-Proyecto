package main

import (
	"github.com/kataras/iris/v12"
	"log"
	"strconv"
	"github.com/pabloalfaro/Car-finder/src/usuario"
	"github.com/pabloalfaro/Car-finder/src/coche"
	"github.com/pabloalfaro/Car-finder/src/anuncio"
	"github.com/pabloalfaro/Car-finder/src/controlador"
)

var	cont = controlador.NuevoControlador()



func newApp() *iris.Application {
    	app := iris.New()
	
	usuario := app.Party("/usuario")
	{
		usuario.Post("", nuevoUsuario)
		usuario.Get("/{user}", buscarUsuario)
		usuario.Put("", modificarUsuario)
		usuario.Delete("/{user}", borrarUsuario)
	}
	
	coche := app.Party("/coche")
	{
		coche.Post("", nuevoCoche)
		coche.Get("", buscarCoche)
		coche.Delete("", borrarCoche)
	}
	
	anuncio := app.Party("/anuncio")
	{
		anuncio.Post("", nuevoAnuncio)
		anuncio.Get("", buscarAnuncio)
		anuncio.Delete("", borrarAnuncio)
	}
	
    return app
}



func nuevoUsuario(ctx iris.Context){
    if ctx.FormValue("username") == ""{
    	log.Print("Valor de username nulo.\n")
    	log.Print("No se ha registrado el usuario.\n")
    	ctx.StatusCode(400)
    	return
    }
	u := ctx.FormValue("username") 
    
	if ctx.FormValue("nombre") == ""{
    	log.Print("Valor de nombre nulo.\n")
    	log.Print("No se ha registrado el usuario.\n")
    	ctx.StatusCode(400)
    	return
    }
	n := ctx.FormValue("nombre")    

	if ctx.FormValue("apellido") == ""{
    	log.Print("Valor de apellido nulo.\n")
    	log.Print("No se ha registrado el usuario.\n")
    	ctx.StatusCode(400)
    	return
    }
	a := ctx.FormValue("apellido")

	if ctx.FormValue("correo") == ""{
    	log.Print("Valor de correo nulo.\n")
    	log.Print("No se ha registrado el usuario.\n")
    	ctx.StatusCode(400)
    	return
    }
	co := ctx.FormValue("correo")
	
	if ctx.FormValue("ciudad") == ""{
    	log.Print("Valor de ciudad nulo.\n")
    	log.Print("No se ha registrado el usuario.\n")
    	ctx.StatusCode(400)
    	return
    }
	ci := ctx.FormValue("ciudad")

	res := cont.NuevoUsuario(u , n , a , co , ci)
	
	if !res{
    	log.Print("Username ya registrado.\n")
    	log.Print("No se ha registrado el usuario.\n")
    	ctx.StatusCode(404)
    	return
	}
	
	ctx.StatusCode(201)
	
	log.Print("Se ha registrado el usuario con éxito.\n")
	
	return
}



func buscarUsuario(ctx iris.Context){
	nombre := ctx.Params().Get("user")
	
	user, res := cont.BuscarUsuario(nombre)
	if res{
		log.Print("Se ha encontrado el usuario.\n")
		ctx.StatusCode(200)
		ctx.JSON(iris.Map{
			"username": usuario.GetUsername(user),
			"nombre": usuario.GetNombre(user),
			"apellido": usuario.GetApellidos(user),
			"correo": usuario.GetCorreo(user),
			"ciudad": usuario.GetCiudad(user),
		})
		return
	}else{
		log.Print("No se ha encontrado el usuario.\n")
		ctx.StatusCode(404)
		return
	}
}



func modificarUsuario(ctx iris.Context){
	u := ctx.FormValue("username")
	
	_, res := cont.BuscarUsuario(u)
	if res{
		log.Print("Se ha encontrado el usuario.\n")
		n := ctx.FormValue("nombre")
		a := ctx.FormValue("apellido")
		co := ctx.FormValue("correo")
		ci := ctx.FormValue("ciudad")
		
		_ = cont.ModificarUsuario(u, n, a, co, ci)
		log.Print("Se ha modificado el usuario correctamente.\n")
		ctx.StatusCode(201)
		return
	}else{
		log.Print("No se ha encontrado el usuario.\n")
		ctx.StatusCode(404)
		return
	}
}



func borrarUsuario(ctx iris.Context){
	nombre := ctx.Params().Get("user")
	
	res := cont.BorrarUsuario(nombre)
	
	if res{
		log.Print("Se ha encontrado el usuario.\n")
		log.Print("Se ha eliminado el usuario correctamente.\n")
		ctx.StatusCode(200)
		return
	}else{
		log.Print("No se ha encontrado el usuario.\n")
		ctx.StatusCode(404)
		return
	}
}



func nuevoCoche(ctx iris.Context){
    if ctx.FormValue("marca") == ""{
    	log.Print("Valor de marca nulo.\n")
    	log.Print("No se ha registrado el coche.\n")
    	ctx.StatusCode(400)
    	return
    }
	ma := ctx.FormValue("marca") 
    
	if ctx.FormValue("modelo") == ""{
    	log.Print("Valor de modelo nulo.\n")
    	log.Print("No se ha registrado el coche.\n")
    	ctx.StatusCode(400)
    	return
    }
	mo := ctx.FormValue("modelo")    

	if ctx.FormValue("serie") == ""{
    	log.Print("Valor de serie nulo.\n")
    	log.Print("No se ha registrado el coche.\n")
    	ctx.StatusCode(400)
    	return
    }
	se := ctx.FormValue("serie")

	if ctx.FormValue("potencia") == ""{
    	log.Print("Valor de potencia nulo.\n")
    	log.Print("No se ha registrado el coche.\n")
    	ctx.StatusCode(400)
    	return
    }
	po, err:= strconv.Atoi(ctx.FormValue("potencia"))
	if err != nil{
		log.Print("Valor de potencia incorrecto.\n")
    	log.Print("No se ha registrado el coche.\n")
    	ctx.StatusCode(400)
    	return
	}
	
	
	res := cont.NuevoCoche(ma, mo, se, po)
	if !res{
    	log.Print("Coche ya registrado.\n")
    	log.Print("No se ha registrado el coche.\n")
    	ctx.StatusCode(404)
    	return
	}
	
	ctx.StatusCode(201)
	
	log.Print("Se ha registrado el coche con éxito.\n")
	return
}



func buscarCoche(ctx iris.Context){
	if ctx.FormValue("marca") == ""{
    	log.Print("Valor de marca nulo.\n")
    	log.Print("No se ha registrado el coche.\n")
    	ctx.StatusCode(400)
    	return
    }
	ma := ctx.FormValue("marca") 
    
	if ctx.FormValue("modelo") == ""{
    	log.Print("Valor de modelo nulo.\n")
    	log.Print("No se ha registrado el coche.\n")
    	ctx.StatusCode(400)
    	return
    }
	mo := ctx.FormValue("modelo")    

	if ctx.FormValue("serie") == ""{
    	log.Print("Valor de serie nulo.\n")
    	log.Print("No se ha registrado el coche.\n")
    	ctx.StatusCode(400)
    	return
    }
	se := ctx.FormValue("serie")

	if ctx.FormValue("potencia") == ""{
    	log.Print("Valor de potencia nulo.\n")
    	log.Print("No se ha registrado el coche.\n")
    	ctx.StatusCode(400)
    	return
    }
	po, err:= strconv.Atoi(ctx.FormValue("potencia"))
	if err != nil{
		log.Print("Valor de potencia incorrecto.\n")
    	log.Print("No se ha registrado el coche.\n")
    	ctx.StatusCode(400)
    	return
	}
	
	car, res := cont.BuscarCoche(ma, mo, se, po)
	if res{
		log.Print("Se ha encontrado el coche.\n")
		ctx.StatusCode(200)
		ctx.JSON(iris.Map{
			"id": coche.GetId(car),
			"marca": coche.GetMarca(car),
			"modelo": coche.GetModelo(car),
			"serie": coche.GetSerie(car),
			"potencia": coche.GetPotencia(car),
		})
		return
	}else{
		log.Print("No se ha encontrado el coche.\n")
		ctx.StatusCode(404)
		return
	}
}



func borrarCoche(ctx iris.Context){
	id, err:= strconv.Atoi(ctx.FormValue("id"))
	if err != nil{
		log.Print("Valor de id incorrecto.\n")
    	log.Print("No se ha registrado el coche.\n")
    	ctx.StatusCode(400)
    	return
	}
	
	res := cont.BorrarCoche(id)
	if res{
		log.Print("Se ha encontrado el coche.\n")
		log.Print("Se ha eliminado el coche correctamente.")
		
		ctx.StatusCode(200)
		return
	}else{
		log.Print("No se ha encontrado el coche.\n")
		ctx.StatusCode(404)
		return
	}
	
}



func nuevoAnuncio(ctx iris.Context){
	if ctx.FormValue("usuario") == ""{
    	log.Print("Valor de usuario nulo.\n")
    	log.Print("No se ha registrado el anuncio.\n")
    	ctx.StatusCode(400)
    	return
    }
    u := ctx.FormValue("usuario")
    
    if ctx.FormValue("precio") == ""{
    	log.Print("Valor de precio nulo.\n")
    	log.Print("No se ha registrado el anuncio.\n")
    	ctx.StatusCode(400)
    	return
    }
    pre, err := strconv.ParseFloat(ctx.FormValue("precio"), 32)
    if err != nil{
		log.Print("Valor de precio incorrecto.\n")
    	log.Print("No se ha registrado el anuncio.\n")
    	ctx.StatusCode(400)
    	return
	}
	preF := float32(pre)   

	if ctx.FormValue("coche") == ""{
    	log.Print("Valor de coche nulo.\n")
    	log.Print("No se ha registrado el anuncio.\n")
    	ctx.StatusCode(400)
    	return
    }
	car, err:= strconv.Atoi(ctx.FormValue("coche"))
	if err != nil{
		log.Print("Valor de coche incorrecto.\n")
    	log.Print("No se ha registrado el anuncio.\n")
    	ctx.StatusCode(400)
    	return
	}
	
	if ctx.FormValue("km") == ""{
    	log.Print("Valor de km nulo.\n")
    	log.Print("No se ha registrado el anuncio.\n")
    	ctx.StatusCode(400)
    	return
    }
	km, err:= strconv.Atoi(ctx.FormValue("km"))
	if err != nil{
		log.Print("Valor de km incorrecto.\n")
    	log.Print("No se ha registrado el anuncio.\n")
    	ctx.StatusCode(400)
    	return
	}
	
	if ctx.FormValue("estado") == ""{
    	log.Print("Valor de estado nulo.\n")
    	log.Print("No se ha registrado el anuncio.\n")
    	ctx.StatusCode(400)
    	return
    }
	es := ctx.FormValue("estado")
	
	if ctx.FormValue("ciudad") == ""{
    	log.Print("Valor de ciudad nulo.\n")
    	log.Print("No se ha registrado el anuncio.\n")
    	ctx.StatusCode(400)
    	return
    }
	ci := ctx.FormValue("ciudad")
	
	if ctx.FormValue("descripcion") == ""{
    	log.Print("Valor de descripcion nulo.\n")
    	log.Print("No se ha registrado el anuncio.\n")
    	ctx.StatusCode(400)
    	return
    }
	des := ctx.FormValue("descripcion")
	
	if ctx.FormValue("color") == ""{
    	log.Print("Valor de color nulo.\n")
    	log.Print("No se ha registrado el anuncio.\n")
    	ctx.StatusCode(400)
    	return
    }
	col := ctx.FormValue("color")
	
	r := cont.NuevoAnuncio(u, preF, car, km, es, ci, des, col)
	
	if r {
		ctx.StatusCode(201)
		
		log.Print("Se ha registrado el anuncio con éxito.\n")
		return
	}else{
		ctx.StatusCode(400)
		log.Print("El usuario no estaba registrado.\n")
		log.Print("No se ha registrado el anuncio con éxito.\n")
		return
	}
}



func buscarAnuncio(ctx iris.Context){
	if ctx.FormValue("usuario") == ""{
    	log.Print("Valor de usuario nulo.\n")
    	log.Print("No se ha registrado el anuncio.\n")
    	ctx.StatusCode(400)
    	return
    }
    u := ctx.FormValue("usuario")
    
    if ctx.FormValue("precio") == ""{
    	log.Print("Valor de precio nulo.\n")
    	log.Print("No se ha registrado el anuncio.\n")
    	ctx.StatusCode(400)
    	return
    }
    pre, err := strconv.ParseFloat(ctx.FormValue("precio"), 32)
    if err != nil{
		log.Print("Valor de precio incorrecto.\n")
    	log.Print("No se ha registrado el anuncio.\n")
    	ctx.StatusCode(400)
    	return
	}
	preF := float32(pre)
	
	car, err:= strconv.Atoi(ctx.FormValue("coche"))
	if err != nil{
		log.Print("Valor de coche incorrecto.\n")
    	log.Print("No se ha registrado el anuncio.\n")
    	ctx.StatusCode(400)
    	return
	}
	
	if ctx.FormValue("km") == ""{
    	log.Print("Valor de km nulo.\n")
    	log.Print("No se ha registrado el anuncio.\n")
    	ctx.StatusCode(400)
    	return
    }
	km, err:= strconv.Atoi(ctx.FormValue("km"))
	if err != nil{
		log.Print("Valor de km incorrecto.\n")
    	log.Print("No se ha registrado el anuncio.\n")
    	ctx.StatusCode(400)
    	return
	}
	
	if ctx.FormValue("estado") == ""{
    	log.Print("Valor de estado nulo.\n")
    	log.Print("No se ha registrado el anuncio.\n")
    	ctx.StatusCode(400)
    	return
    }
	es := ctx.FormValue("estado")
	
	if ctx.FormValue("ciudad") == ""{
    	log.Print("Valor de ciudad nulo.\n")
    	log.Print("No se ha registrado el anuncio.\n")
    	ctx.StatusCode(400)
    	return
    }
	ci := ctx.FormValue("ciudad")
	
	if ctx.FormValue("descripcion") == ""{
    	log.Print("Valor de descripcion nulo.\n")
    	log.Print("No se ha registrado el anuncio.\n")
    	ctx.StatusCode(400)
    	return
    }
	des := ctx.FormValue("descripcion")
	
	if ctx.FormValue("color") == ""{
    	log.Print("Valor de color nulo.\n")
    	log.Print("No se ha registrado el anuncio.\n")
    	ctx.StatusCode(400)
    	return
    }
	col := ctx.FormValue("color")
	
	
	index, res := cont.BuscarAnuncio(u, preF, car, km, es, ci, des, col)
	
	if res{
		log.Print("Se ha encontrado el coche.\n")
		ctx.StatusCode(200)
		ctx.JSON(iris.Map{
			"id": anuncio.GetId(index),
			"usuario": anuncio.GetUsuario(index),
			"precio": anuncio.GetPrecio(index),
			"coche": anuncio.GetCoche(index),
			"kilometros": anuncio.GetKm(index),
			"estado": anuncio.GetEstado(index),
			"ciudad": anuncio.GetCiudad(index),
			"descripcion": anuncio.GetDescripcion(index),
			"color": anuncio.GetColor(index),
		})
		return
	}else{
		log.Print("No se ha encontrado el anuncio.\n")
		ctx.StatusCode(404)
		return
	}
}



func borrarAnuncio(ctx iris.Context){
	if ctx.FormValue("id") == ""{
    	log.Print("Valor de id nulo.\n")
    	log.Print("No se ha registrado el anuncio.\n")
    	ctx.StatusCode(400)
    	return
    }
	id, err:= strconv.Atoi(ctx.FormValue("id"))
	if err != nil{
		log.Print("Valor de id incorrecto.\n")
    	log.Print("No se ha registrado el anuncio.\n")
    	ctx.StatusCode(400)
    	return
	}
	
	res := cont.BorrarAnuncio(id)
	if res{
		log.Print("Se ha encontrado el anuncio.\n")
		log.Print("Se ha eliminado el anuncio correctamente.")
		
		ctx.StatusCode(200)
		return
	}else{
		log.Print("No se ha encontrado el anuncio.\n")
	
		ctx.StatusCode(404)
		return
	}
}



func main() {
	app := newApp()
	app.Listen(":8080")
}
