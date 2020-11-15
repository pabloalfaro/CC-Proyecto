package controlador_test

import (
	. "github.com/pabloalfaro/Car-finder/src/controlador"
	. "github.com/pabloalfaro/Car-finder/src/coche"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Controlador", func() {
	var (
		cont1 Controlador
		cont2 Controlador
		coche Coche
	)
	
	BeforeEach(func() {
		cont1 = NuevoControlador()
		
		cont2 = NuevoControlador()
		_ = cont2.NuevoUsuario("u", "n", "a", "co", "ci")
		
		coche = NewCoche("ma", "mo", "s", 90)
	})
	
	Describe("Un nuevo usuario quiere registrarse en la aplicación.", func() {
		Context("Se quiere registrar con un nombre de usuario que nadie ha utilizado.", func() {
			It("Debería poder registrarse.", func(){
				Expect(cont1.NuevoUsuario("u", "n", "a", "co", "ci")).To(BeTrue())
			})
		})
		
		Context("El nombre de usuario esta ya registrado.", func() {
			It("No debería poder registrarse.", func(){
				Expect(cont2.NuevoUsuario("u", "n", "a", "co", "ci")).To(BeFalse())
			})
		})
	})
	
	Describe("Un usuario quiere subir un anuncio.", func() {
		Context("El usuario se ha registrado correctamente de manera previa.", func() {
			It("Debería poder subir el anuncio.", func(){
				Expect(cont2.NuevoAnuncio("u" , 1000, coche, 100000, "e", "ciu", "d", "col")).To(BeTrue())
			})
		})
		
		Context("El usuario no se ha registrado correctamente de manera previa.", func() {
			It("No debería poder subir el anuncio.", func(){
				Expect(cont1.NuevoAnuncio("u" , 1000, coche, 100000, "e", "ciu", "d", "col")).To(BeFalse())
			})
		})
	})
	
	
})
