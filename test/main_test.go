package test

import "testing"
import "github.com/pabloalfaro/Car-finder/src/controlador"


//Hacemos una inserción posible
func Test_nuevoUsuario1(t *testing.T){
	t.Log("Test 1: inserción correcta de usuario.")	
	
	//Inicio el map de Usuarios
	controlador.Main()
	
	b := controlador.NuevoUsuario("u", "n", "a", "co", "ci")
	
	if b {
		t.Log("Test 1. OK. Inserción válida.")
	}else{
		t.Error("No es válida la inserción de usuario.")
	}
}


//Hacemos una inserción que no es posible
func Test_nuevoUsuario2(t *testing.T){
	t.Log("Test 2: inserción imposible de usuario.")	
	
	//Inicio el map de Usuarios
	controlador.Main()
	
	_ = controlador.NuevoUsuario("u", "n", "a", "co", "ci")
	b := controlador.NuevoUsuario("u", "n", "a", "co", "ci")
	
	if b {
		t.Error("Permite insertar dos vecesel mismo usuario.")

	}else{
		t.Log("Test 2. OK. No se ha permitido la inserción.")		
	}
}
