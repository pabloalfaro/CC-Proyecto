[![Build Status](https://travis-ci.com/pabloalfaro/Car-finder.svg?branch=main)](https://travis-ci.com/pabloalfaro/Car-finder)

# Car finder (En desarrollo)
Soy Pablo Alfaro, alumno de la asignatura de Cloud Computing del Máster en Ingeniería Informática de la Universidad de Granada. En este repositorio se encuentra el proyecto de la asignatura.


## Gestor de tareas (Entrega 2)

Para mi proyecto he utilizado [Makefile](https://github.com/pabloalfaro/Car-finder/blob/main/Makefile). Me he decido por este gestor porque es uno de los más utilizados. Cualquiera que se encuentre con este tipo de ficheros sabe como uitilizar las instrucciones básicas sin tener que buscar información al respecto. He implementado una funcionalidad de test, una funcionalidad "fmt" que hace que el código programado sea más legible, la de install y un all. Para escribir mi fichero he seguido algunas buenas prácticas para Golang descritas en esta [Web](https://ops.tips/blog/minimal-golang-makefile/).

## Biblioteca de aserciones (Entrega 2)

La biblioteca de aserciones que he utilizado ha sido [Gomega](https://github.com/onsi/gomega). En este caso la elección ha venido dada por ser el recomendado por mi marco de pruebas. En mi caso he utilizado aserciones para comprobar las salidas de mis métodos utilizando la funcion `Expect()`, `To()`, `BeTrue()` y `BeFalse()`. Estos métodos devuelven un valor de True cuando se ha completado con éxito y un valor de False en el caso contrario.

## Marco de pruebas (Entrega 2)

Me he decidido por tests de tipo bdd, este tipo de test se han de implementar para comprobar que se cumple con las funcionalidades descritas en las distintas HU. En este caso se hace de una manera casi literal, intentando que los tests se parezcan lo máximo posible al lenguaje natural. He hecho pruebas con [Godog](https://github.com/cucumber/godog) y con [Ginkgo](https://github.com/onsi/ginkgo). He probado ambos frameworks, pero con el de Godog no conseguí correr los tests, las implementaciones de pueden ver en el commit [#47](https://github.com/pabloalfaro/Car-finder/commit/afea01ddabb865584b1fa0807eed2b5aacf6453f). Después de intentarlo con Godog probé con Ginko, que es con el que finalmente me he quedado. Tras haber usado los dos y aunque hubiese podido correr los de Godog, me hubiese quedado con Ginkgo. Me ha resultado menos laborioso implementar los test y personalmente veo que el código es más claro en este caso. 

## Ejecución de los tests (Entrega 2)

Ginkgo tiene una manera pripia de ejecutar los tests pero se puede hacer con la herramienta de Go, utilizando el comando `go test`. De esta manera ejecutas los test que tengas en el directorio dónde lo estés ejecutando. Otra método de ejecutrarlos sería con `go test ./ruta_de_tests`, de esta forma puedes elegir el directorio dónde están los test que quieras ejecutar. Los archivos de Go de test se identifican porque tienen este tipo de nombres: `nombre_test.go`. Estos archivos puedes guardarlos junto con la clase que testean o los puedes guardar en otro directorio a parte. Go no permite normalmente que tengas ficheros de dos `package` distintos, pero con los tests hace una excepción. En mi caso he guardado los test junto con la clase a la que hacen referencia.

## Elección de imagen para Docker (Entrega 3)

Quiero crear una imagen Docker para ejecutar los test de mi aplicación. El objetivo es que esta imagen tenga instalado Golang, en mi caso la versión 1.15.5. También quiero que la imagen ocupe lo menos posible. Para ello, voy a repasar algunas imágenes de algunos sistemas operativos y después las versiones oficiales. Tras esto también revisaré documentación para informarme sobre la imagen más extendida.

Para hacer las comparaciones de las imágenes me he probado dos herramientas:

-	[Docker-Diff](https://github.com/moul/docker-diff)

-	[Container-Diff](https://github.com/GoogleContainerTools/container-diff)

Después de probar ambas me he quedado con Container-Diff, ya que su uso me ha resultado más sencillo y el repositorio estaba más actualizado. Utilizando esta herramienta he podido ver el tamaño de las imágenes y los ficheros que tenían, entre otras cosas.

Revisé la versión latest de los SO Fedora, CentOS, Alpine y Ubuntu. Comparándolos he visto, como era de esperar, que la versión de Alpine era la que menos ocupaba. Revisando las 4 imágenes he visto que no tenían ninguna la instalación de Go por lo que el tamaño final será mayor.
Las imágenes oficiales que he probado han sido las de Buster y Alpine. De nuevo la de Alpine era la de menor tamaño y en este caso la de Buster era la de mayor tamaño de todas las de la prueba. 

|Imagen        |	Peso  | Golang |
|--------------|--------|--------|
|Fedora:33     | 166.9M | NO     |
|Centos:centos8| 205.1M |	NO     |
|Alpine: 3.12.1| 5.3M	  | NO     |
|Ubuntu: 20.04 | 72.9M	| NO     |
|Golang:buster | 797.5M |	SI     |
|Golang:alpine | 284.7M |	SI     |

Tras comparar estas imágenes he pasado a revisar documentación sobre Dockers para Golang :

- https://www.docker.com/blog/containerize-your-go-developer-environment-part-1/

- https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324

-	https://levelup.gitconnected.com/complete-guide-to-create-docker-container-for-your-golang-application-80f3fb59a15e

En esta documentación y alguna más que he visto he comprobado que la imagen más usada es la de golang:alpine. Dado que era la mejor de las que había visto y es la más extendida, será mi elección para construir el Docker de pruebas.

## Ejecución de los test (Entrega 3)

Para hacer la build del docker he usado el sioguiente comando:

`docker build --no-cache -t golang_alpine -f Dockerfile .`

A mi imagen le he llamado golang_alpine y el fichero que he utilizado ha sido [Dockerfile](https://github.com/pabloalfaro/Car-finder/blob/main/Dockerfile). Con la imagen construida lanzo los test de la siguiente manera:

`docker run -t -v pwp:/app/test golang_alpine:latest`

Así lanzo los test con la imagen local. Se pueden ejecutar también con la imagen remota de mi repositorio de Docker Hub. Mi usuario es pabloalfaro y mi repo es car-finder. Se podría usar la imagen remota de esta forma: (utilizo la etiqueta latest para referirme a la ultima imagen construida)

`docker run -t -v pwp:/app/test pabloalfaro/car-finder:latest`

En ambos casos trabajo sobre el directorio /app/test, es el que he preparado en Dockerfile para trabajar. El resultado que obtengo con las ejecuciones se puede ver en la imagen de abajo.

![test](https://github.com/pabloalfaro/Car-finder/blob/main/Documentaci%C3%B3n%20adicional/run%20en%20docker.png)

## Trabajo con Docker Hub (Entrega 3)

Lo primero que he hecho ha sido registrarme en la página y crear mi [repo](https://hub.docker.com/r/pabloalfaro/car-finder) para el proyecto. En primer lugar he querido hacer un push de una imagen. Para ello he seguido los pasos de este [tutorial](https://ropenscilabs.github.io/r-docker-tutorial/04-Dockerhub.html). Lo primero que he hecho ha sido el login:

`docker login --username=pabloalfaro`

Una vez identificado he creado un tag para hacer el push:

`docker tag 8b308631375e pabloalfaro/car-finder:firsttry`

Con este tag, firsttry, creado he pasado a hacer el push:

`docker push pabloalfaro/car-finder`

El siguente paso ha sido automatizar las builds cada vez que hiciese un commit en mi repo de GitHub. Esto se puede hacer desde Manage Repository, Builds, Configure Automated Builds. Hay que enlazar tu cuenta con la de GitHub, una vez hecho, tienes que elegir el repo, la rama y la localización del fichero Dockerfile. 

![configuración](https://github.com/pabloalfaro/Car-finder/blob/main/Documentaci%C3%B3n%20adicional/configuracion%20DockerHub.png)

Para terminar puedes guardar la configuración usando *Save* o puedes hacer el primer build con *Save and Build*. Cuando ya lo tienes configurado puedes ver las builds así:

![builds](https://github.com/pabloalfaro/Car-finder/blob/main/Documentaci%C3%B3n%20adicional/builds.png)

## Historias de usuario

En mi caso he pensado en 3 posibles usuarios, estos son:

- El cliente que ha encargado el proyecto: [Cliente](https://github.com/pabloalfaro/Car-finder/issues?q=is%3Aissue+is%3Aopen+label%3Acliente)
- El comprador de los coches: [Comprador](https://github.com/pabloalfaro/Car-finder/issues?q=is%3Aissue+is%3Aopen+label%3Acomprador)
- El vendedor de los coches: [Vendedor](https://github.com/pabloalfaro/Car-finder/issues?q=is%3Aissue+is%3Aopen+label%3Avendedor)


## Documentación adicional
- [Roadmap del proyecto](https://github.com/pabloalfaro/Car-finder/blob/main/roadmap.md)
- [Clases programadas](https://github.com/pabloalfaro/Car-finder/tree/main/src)
- [Tests unitarios](https://github.com/pabloalfaro/Car-finder/tree/main/src/controlador)
- [PMV 1](https://github.com/pabloalfaro/Car-finder/milestone/3)
- [PMV 2](https://github.com/pabloalfaro/Car-finder/milestone/4)
- [Correciones para la reentrega](https://github.com/pabloalfaro/Car-finder/milestone/5)


## Contenido anterior
Aquí dejo los links a los contenidos de las semanas anteriores:

- [Tema 1](https://github.com/pabloalfaro/Car-finder/blob/main/Semanas%20anteriores/tema1.md)
- [Tema 2](https://github.com/pabloalfaro/Car-finder/blob/main/Semanas%20anteriores/tema2.md)
