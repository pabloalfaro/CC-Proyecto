Feature: controlador 

  Un usario debe poder registrarse con un nombre de usuario válido y ha de poder registrar anuncios.
  
 Scenario: Creación de usuario válido
   Given no hay usuarios registrados con mi username
   When me registro introduciendo datos válidos
   Then tengo mi usuario registrado
   
 Scenario: Creación de usuario no válido
   Given hay usuarios registrados con mi username
   When me registro introduciendo datos que no son válidos
   Then no me deja registrar mi usuario
 
