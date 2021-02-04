# Test de la composición

## Resumen introductorio

El objetivo ha sido probar el correcto funcionamiento de la composición que había desarrollado. Esta composición la forman una api y una base de datos Mysql. La forma de 
probar esta composición ha sido utilizando un cliente que, tras haber levantado la imagen, hace una serie de peticiones a la api que a su vez las traslada a la base de datos.
En un principio este cliente únicamente se conectaba a la api, que debía estar operativa en el momento en el que se quisiesen hacer los test. Para que la prueba se pudiese 
realizar sin necesidad de haber iniciado la api previamente de manera externa, he programado el cliente para que inicie la composición al inicio y la pare al terminar de hacer la 
última petición. De esta manera para probar la composición basta con ejecutar [cliente.go](https://github.com/pabloalfaro/Car-finder/blob/main/src/cliente/cliente.go).
A continuación, voy a mostrar cuáles has sido las historias de usuario testeadas y las respuestas a las peticiones. Se puede observar que los códigos de respuesta nos indican que 
se están llevando a cabo correctamente las operaciones.

## Funcionalidades testeadas de usuarios
[[HU]Quiero poder gestionar los usuarios de la página](https://github.com/pabloalfaro/Car-finder/issues/7)

![usuarios](https://github.com/pabloalfaro/Car-finder/blob/main/docs/Documentaci%C3%B3n%20adicional/usuario.png)

## Funcionalidades testeadas de coches
[[HU]Quiero poder gestionar los coches de la página](https://github.com/pabloalfaro/Car-finder/issues/71)

![coches](https://github.com/pabloalfaro/Car-finder/blob/main/docs/Documentaci%C3%B3n%20adicional/coche.png)

## Funcionalidades testeadas de anuncios
[[HU]Quiero poder anunciar los coches](https://github.com/pabloalfaro/Car-finder/issues/8)

[[HU]Quiero poder hacer búsquedas entre los distintos anuncios](https://github.com/pabloalfaro/Car-finder/issues/6)

[[HU]Quiero poder eliminar mis anuncios](https://github.com/pabloalfaro/Car-finder/issues/10)

![anuncios](https://github.com/pabloalfaro/Car-finder/blob/main/docs/Documentaci%C3%B3n%20adicional/anuncio.png)

### Documentación
He obtenido información para desarrollar este cliente de las siguientes páginas:

Cómo hacer peticiones POST, GET, PUT y DELETE: https://parzibyte.me/blog/2019/05/21/peticion-post-get-put-delete-go-net-http/

Cómo pausar la ejecución (así dejar tiempo a cargar la composición al inicio): https://golangcode.com/sleeping-with-go/

Me he servido de la herramienta [postman](https://www.postman.com/) para hacer algunas pruebas antes de tener el cliente terminado. Resulta muy útil para realizar peticiones a
una api.
