package controlador

import (
	"fmt"
	"github.com/cucumber/godog"
)

var testControlador *controlador

func sinUsuarioPrevioRegistrado() error{
	testControlador.usuarios = make(map[string]usuario.Usuario)
	
	return nil
}

func nuevoUsuario(u string, n string, a string, co string, ci string) error{
	testControlador.NuevoUsuario(u string, n string, a string, co string, ci string)
	
	return nil
}

func hayUnUsuarioNuevo(u string) error{
	_, ok := testControlador.usuarios[u]
	
	if ok{
		return(nil)
	}
	
	return fmt.error("No se ha podido introducir el nuevo usuario")
}

func FeatureContext(s
