## Test de la api

Al igual que el resto del código, hay que testear la api. Para realizar estos test he usado la librería de `GO` [`testing`](https://golang.org/pkg/testing/) y una librería para 
poder hacer las peticiones llamada [`httptest`](https://golang.org/pkg/net/http/httptest/).

Los bloques que he testeado han sido los siguientes:
- Usuarios
- Coches 
- Anuncios 

### Funcionalidades testeadas de usuarios
[[HU]Quiero poder gestionar los usuarios de la página](https://github.com/pabloalfaro/Car-finder/issues/7)

### Funcionalidades testeadas de coches
[[HU]Quiero poder gestionar los coches de la página](https://github.com/pabloalfaro/Car-finder/issues/71)

### Funcionalidades testeadas de anuncios
[[HU]Quiero poder anunciar los coches](https://github.com/pabloalfaro/Car-finder/issues/8), 
[[HU]Quiero poder hacer búsquedas entre los distintos anuncios](https://github.com/pabloalfaro/Car-finder/issues/6) y
[[HU]Quiero poder eliminar mis anuncios](https://github.com/pabloalfaro/Car-finder/issues/10).

En este apartado he visto que la funcionalidad de buscar anuncio. Los test de `buscarAnuncio` y `borrarAnuncio` los he dejado comentados para que no diesen error. Dejo esto 
reflejado en este [issue](https://github.com/pabloalfaro/Car-finder/issues/72).

## Implementación
Muestro a continuación la implementación de los test de usuarios, que son los más completos.

~~~
func TestUsuario(t *testing.T) {
    app := newApp()
    e := httptest.New(t, app)
    
    log.Print("TEST 1: CREAR USUARIO.")
 
    e.POST("/usuario").
    	WithFormField("username", "PabloA").
    	WithFormField("nombre", "Pablo").
    	WithFormField("apellido", "Alfaro").
    	WithFormField("correo", "ejemplo@correo.com").
    	WithFormField("ciudad", "Pamplona").
    	Expect().Status(201)
    	
    
    log.Print("TEST 2: BUSCAR USUARIO.")
    	
    r:= e.GET("/usuario/buscar/{user}", "PabloA").
    		Expect().Status(200).
    		JSON().Object()
    		
    r.Value("username").Equal("PabloA")
    r.Value("nombre").Equal("Pablo")
    r.Value("apellido").Equal("Alfaro")
    r.Value("correo").Equal("ejemplo@correo.com")
    r.Value("ciudad").Equal("Pamplona")
    
    
    log.Print("TEST 3: MODIFICAR USUARIO.")
    
    e.PUT("/usuario/modificar").
    	WithFormField("username", "PabloA").
    	WithFormField("nombre", "Pablo").
    	WithFormField("apellido", "Alfaro").
    	WithFormField("correo", "ejemplo@correo.com").
    	WithFormField("ciudad", "Pamplona").
    	Expect().Status(201)
    	
    	
    log.Print("TEST 4: BORRAR USUARIO.")
    
    e.DELETE("/usuario/borrar/{user}", "PabloA").
    		Expect().Status(200)
    
}
~~~

Se puede ver el documento completo en [test](https://github.com/pabloalfaro/Car-finder/blob/main/src/api/api_test.go).
