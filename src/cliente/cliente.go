package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"bytes"
	"mime/multipart"
	"fmt"
)

func main(){
	url_u1 := "http://localhost:8080/usuario"
	url_u2 := "http://localhost:8080/usuario/PabloG"
	
	url_c1 := "http://localhost:8080/coche"
	url_c2 := "http://localhost:8080/coche/2"
	
	url_a1 := "http://localhost:8080/anuncio"
	url_a2 := "http://localhost:8080/anuncio/1"
	
	
 	payload_u := &bytes.Buffer{}
	writer_u := multipart.NewWriter(payload_u)
	_ = writer_u.WriteField("username", "PabloG")
	_ = writer_u.WriteField("nombre", "Pablo")
	_ = writer_u.WriteField("apellido", "Alfaro")
	_ = writer_u.WriteField("correo", "pabloalfaro@correo.ugr.es")
	_ = writer_u.WriteField("ciudad", "Pamplona")
	err := writer_u.Close()
	if err != nil {
	  log.Print(err)
	  return
	}
  
	log.Print("---Añadir usuario---")
	peticion(url_u1, "POST", payload_u, writer_u)
	log.Print("---Buscar usuario---")
	peticion(url_u2, "GET", payload_u, writer_u)
	log.Print("---Actualizar usuario---")
	peticion(url_u1, "PUT", payload_u, writer_u)
	log.Print("---Borrar usuario---")
	peticion(url_u2, "DELETE", payload_u, writer_u)
	
	
	payload_c := &bytes.Buffer{}
  writer_c := multipart.NewWriter(payload_c)
  _ = writer_c.WriteField("marca", "Audi")
  _ = writer_c.WriteField("modelo", "A3")
  _ = writer_c.WriteField("serie", "S")
  _ = writer_c.WriteField("potencia", "180")
  _ = writer_c.WriteField("id", "1")
  err = writer_c.Close()
  if err != nil {
	  log.Print(err)
	  return
	}
  
	log.Print("---Añadir coche---")
	peticion(url_c1, "POST", payload_c, writer_c)
	log.Print("---Buscar coche---")
	peticion(url_c2, "GET", payload_c, writer_c)
	log.Print("---Borrar coche---")
	peticion(url_c2, "DELETE", payload_c, writer_c)
	
	
  payload_a := &bytes.Buffer{}
  writer_a := multipart.NewWriter(payload_a)
  _ = writer_a.WriteField("usuario", "PabloG")
  _ = writer_a.WriteField("precio", "10000")
  _ = writer_a.WriteField("coche", "1")
  _ = writer_a.WriteField("km", "100000")
  _ = writer_a.WriteField("estado", "Correcto")
  _ = writer_a.WriteField("ciudad", "Granada")
  _ = writer_a.WriteField("descripcion", "Audi en buen estado")
  _ = writer_a.WriteField("color", "Negro")
  _ = writer_a.WriteField("id", "2")
  err = writer_a.Close()
  if err != nil {
    fmt.Println(err)
    return
  }
	
	log.Print("---Añadir anuncio---")
	//peticion(url_u1, "POST", payload_u, writer_u)
	peticion(url_a1, "POST", payload_a, writer_a)
	log.Print("---Buscar anuncio---")
	peticion(url_a2, "GET", payload_a, writer_a)
	log.Print("---Borrar anuncio---")
	peticion(url_a2, "DELETE", payload_a, writer_a)
	
}

func peticion (url string, cod string, payload *bytes.Buffer, writer *multipart.Writer){
	clienteHttp := &http.Client{}
	peticion, err := http.NewRequest(cod, url, payload)
	// Compruebo el error al crear la petición
	if err != nil {
		log.Fatalf("Error creando petición: %v", err)
	}
	peticion.Header.Set("Content-Type", writer.FormDataContentType())
	
	respuesta, err := clienteHttp.Do(peticion)
	// Compruebo el error al hacer la petición
	if err != nil {
		log.Fatalf("Error haciendo petición: %v", err)
	}
	
	// Cierro el cuerpo
	defer respuesta.Body.Close()
	
	cuerpoRespuesta, err := ioutil.ReadAll(respuesta.Body)
	// Compruebo el error al leer la respuesta
	if err != nil {
		log.Fatalf("Error leyendo respuesta: %v", err)
	}

	respuestaString := string(cuerpoRespuesta)
	log.Printf("Código de respuesta: %d", respuesta.StatusCode)
	log.Printf("Cuerpo de respuesta del servidor: '%s'", respuestaString)
}

