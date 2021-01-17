package main

import (
    "testing"
    "log"
    "github.com/kataras/iris/v12/httptest"
)

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



func TestCoche(t *testing.T) {
    app := newApp()
    e := httptest.New(t, app)
    
    log.Print("TEST 1: CREAR COCHE.")
 
    e.POST("/coche").
    	WithFormField("marca", "Audi").
    	WithFormField("modelo", "A4").
    	WithFormField("serie", "S").
    	WithFormField("potencia", "180").
    	Expect().Status(201)
    	
    
    log.Print("TEST 2: BUSCAR COCHE.")
    	
    r:= e.GET("/coche/buscar").
		WithQuery("marca", "Audi").
    	WithQuery("modelo", "A4").
    	WithQuery("serie", "S").
    	WithQuery("potencia", "180").
		Expect().Status(200).
		JSON().Object()
    		
    r.Value("marca").Equal("Audi")
    r.Value("modelo").Equal("A4")
    r.Value("serie").Equal("S")
    r.Value("potencia").Equal(180)
    	
    	
    log.Print("TEST 3: BORRAR COCHE.")
    
    e.DELETE("/coche/borrar").
    	WithQuery("marca", "Audi").
    	WithQuery("modelo", "A4").
    	WithQuery("serie", "S").
    	WithQuery("potencia", "180").
    	Expect().Status(200)
    
}


func TestAnuncio(t *testing.T) {
    app := newApp()
    e := httptest.New(t, app)
    
    
    log.Print("TEST 1: CREAR ANUNCIO.")
    
    e.POST("/usuario").
    	WithFormField("username", "PabloA").
    	WithFormField("nombre", "Pablo").
    	WithFormField("apellido", "Alfaro").
    	WithFormField("correo", "ejemplo@correo.com").
    	WithFormField("ciudad", "Pamplona").
    	Expect().Status(201)
    
    e.POST("/anuncio").
    	WithFormField("usuario", "PabloA").
    	WithFormField("precio", "10000").
    	WithFormField("marca", "Audi").
    	WithFormField("modelo", "A4").
    	WithFormField("serie", "S").
    	WithFormField("potencia", "180").
    	WithFormField("km", "50000").
    	WithFormField("estado", "Bueno").
    	WithFormField("ciudad", "Pamplona").
    	WithFormField("descripcion", "PabloA").
    	WithFormField("color", "negro").
    	Expect().Status(201)
    	
    	
    log.Print("TEST 2: BUSCAR ANUNCIO.")
    /*
    r:= e.GET("/anuncio/buscar").
		WithQuery("usuario", "PabloA").
    	WithQuery("precio", "10000").
    	WithQuery("marca", "Audi").
    	WithQuery("modelo", "A4").
    	WithQuery("serie", "S").
    	WithQuery("potencia", "180").
    	WithQuery("km", "50000").
    	WithQuery("estado", "Bueno").
    	WithQuery("ciudad", "Pamplona").
    	WithQuery("descripcion", "Desc").
    	WithQuery("color", "negro").
		Expect().Status(200).
		JSON().Object()
    	
    r.Value("potencia").Equal(180)
    r.Value("usuario").Equal("PabloA")
	r.Value("precio").Equal("10000")
	r.Value("kilometros").Equal(50000)
	r.Value("estado").Equal("Bueno")
	r.Value("ciudad").Equal("Pamplona")
	r.Value("descripcion").Equal("Desc")
	r.Value("color").Equal("negro")
   */
   
   
    log.Print("TEST 3: BORRAR ANUNCIO.")
    /*
    e.DELETE("/anuncio/borrar").
    	WithQuery("usuario", "PabloA").
    	WithQuery("precio", "10000").
    	WithQuery("marca", "Audi").
    	WithQuery("modelo", "A4").
    	WithQuery("serie", "S").
    	WithQuery("potencia", "180").
    	WithQuery("km", "50000").
    	WithQuery("estado", "Bueno").
    	WithQuery("ciudad", "Pamplona").
    	WithQuery("descripcion", "Desc").
    	WithQuery("color", "negro").
    	Expect().Status(200)
    */
}
