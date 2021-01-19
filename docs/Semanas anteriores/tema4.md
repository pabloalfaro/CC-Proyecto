## Elección de imagen para Docker 

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

## Ejecución de los test

Para hacer la build del docker he usado el sioguiente comando:

`docker build --no-cache -t golang_alpine -f Dockerfile .`

A mi imagen le he llamado golang_alpine y el fichero que he utilizado ha sido [Dockerfile](https://github.com/pabloalfaro/Car-finder/blob/main/Dockerfile). Con la imagen construida lanzo los test de la siguiente manera:

`docker run -t -v pwp:/app/test golang_alpine:latest`

Así lanzo los test con la imagen local. Se pueden ejecutar también con la imagen remota de mi repositorio de Docker Hub. Mi usuario es pabloalfaro y mi repo es car-finder. Se podría usar la imagen remota de esta forma: (utilizo la etiqueta latest para referirme a la ultima imagen construida)

`docker run -t -v pwp:/app/test pabloalfaro/car-finder:latest`

En ambos casos trabajo sobre el directorio /app/test, es el que he preparado en Dockerfile para trabajar. El resultado que obtengo con las ejecuciones se puede ver en la imagen de abajo.

![test](https://github.com/pabloalfaro/Car-finder/blob/main/Documentaci%C3%B3n%20adicional/run%20en%20docker.png)

## Trabajo con Docker Hub

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
