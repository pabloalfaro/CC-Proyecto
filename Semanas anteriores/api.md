## Desarrollo de la api

En mi aplicación podemos diferenciar 3 bloques distintos:
- Usuarios
- Coches 
- Anuncios

Cada uno de estos bloques tienen funcionalidades distintas que pasaré a detallar a continuación.

### Usuarios
Estas funcionalidades corresponden a la [[HU]Quiero poder gestionar los usuarios de la página](https://github.com/pabloalfaro/Car-finder/issues/7)

- `nuevoUsuario` en la ruta `/usuario`. Utilizo el método POST para pasar los parámetros necesarios en forma de formulario. Si alguno de los parámetros es nulo, 
devuelve un error 400. Si se está intentando registrar con un `username` ya registrado devolverá un código de error 409. En el caso de que se realice la petición correctamente 
el código devuelto será el 201.

- `buscarUsuario` en la ruta `/usuario/buscar/{user}`. Utilizo el método GET y le paso el parámetro `user` cómo índice para la búsqueda. Si se ha encontrado el usuario se 
devuelven sus datos con el código 200. En el caso de no encontrar un usuario con ese `username` el código será el 404.

- `modificarUsuario` en la ruta `/usuario/modificar`. Utilizo el método PUT para pasar los parámetros nuevos. Se identifica qué parámetros llegan y se actualizan sólo esos campos.
Si se ha actualizado el usuario correctamente obtendremos un código 201, en el caso contrario el código será el 404 al no haber encontrado el usuario que queremos modificar.

- `borrarUsuario` en la ruta `/usuario/borrar/{user}`. Utilizo el método DELETE y le paso el parámetro `user` cómo índice para borrar. Se busca si existe un usuario con un 
`username` igual a el parámetro `user` y si es así se elimina. De llevarse a cabo correctamente obtendremos un código 201, en el caso contrario el código será el 404 al 
no haber encontrado el usuario que queremos borrar.



## Implementación de la api

En este apartado voy a mostrar como implemento la API para la funcionalidad de Usuarios, que es la más completa.

### Definición de las rutas
~~~
func newApp() *iris.Application {
    app := iris.New()
	
	usuario := app.Party("/usuario")
	{
		usuario.Post("", nuevoUsuario)
		usuario.Get("/buscar/{user}", buscarUsuario)
		usuario.Put("/modificar", modificarUsuario)
		usuario.Delete("/borrar/{user}", borrarUsuario)
	}
      return app
}
~~~

~~~
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
    	ctx.StatusCode(409)
    	return
	}
	
	ctx.StatusCode(201)
	
	log.Print("Se ha registrado el usuario con éxito.\n")
	
	return
}
~~~

~~~
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
~~~

~~~
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
~~~

~~~
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
~~~

El resto de la implementación se encuentra en [api.go](https://github.com/pabloalfaro/Car-finder/blob/main/src/api/api.go).
