### Logs 

Una buena práctica que se puede llevar a cabo a la hora de implementar la api es la de establecer una serie de log que nos indiquen que está pasando durante la ejecución. En el 
caso de `GO`, estos logs vienen implementados en una librería del lenguaje llamada [`log`](https://golang.org/pkg/log/). Esta librería es la que he utilizado para documentar las
actividades de la api.

### Implementación

En mi caso he ido documentando los posibles errores que podían ocurrir y la causa de los mismos. También he indicado si se habían llevado a cabo las tareas. A continuación muestro
un fragmento de código de la funcion nuevoUsuario dónde se puede ver que utilizo la función `log.Print()` para indicar errores, los motivos del error y si se ha completado correctamente.

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
