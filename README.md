[![Build Status](https://travis-ci.com/pabloalfaro/Car-finder.svg?branch=main)](https://travis-ci.com/pabloalfaro/Car-finder)
[![Run Status](https://api.shippable.com/projects/5fd349132e187a0006fc2f1d/badge?branch=main)]()

# Car finder (En desarrollo)
Soy Pablo Alfaro, alumno de la asignatura de Cloud Computing del Máster en Ingeniería Informática de la Universidad de Granada. En este repositorio se encuentra el proyecto de la asignatura.

## Integración continua

Previamente había implementado una serie de test para el código que había programado. Existen herramientas que automatizan la ejecución de los test. Estas herramientas hacen que  conforme se avance con el proyecto se compruebe que las funcionalidades previas no tienen errores además de ir comprobrando las nuevas con nuevos tests. [Travis-CI](https://travis-ci.com/) es un sistema de integración continua con un uso muy extendido en GitHub. Tiene una integración directa con GitHub y permite testear y construir apliaciones en una extensa lista de lenguajes de programación.

En primer lugar me he registrado en Travis con mi cuenta de GitHub. Ya registrado he podido elegir en que repos quería activar la integración continua. El siguiente paso es añadir un fichero [.travis.yml](https://github.com/pabloalfaro/Car-finder/blob/main/.travis.yml). He preparado una primera configuración que lanze los tests usando mi fichero Makefile.

~~~
language: go

go:
- 1.15

script: 
- make test
~~~

Tras hacer la prueba y ver que ejecuta los test correctamente, decidí implementar la ejecución de los tests usando el docker que había configurado previamente. En este caso el fichero de configuración de Travis se quedó de la siguiente forma:

~~~
language: go

go:
- 1.15

services:
  - docker

install:
  - docker pull pabloalfaro/car-finder

script:
  - docker run -t -v $TRAVIS_BUILD_DIR:/app/test pabloalfaro/car-finder
~~~

Con Travis ya configurado he probado otro sistema. El que he elegido ha sido [Shippable](https://app.shippable.com/). La implementación de este sistema es muy similar a Travis, hay que registrarse con una cuenta de GitHub y elegir los repos sobre los que quieres ejecutar la integración continua. A diferencia de Travis, Shippable si que deja que te registres con otros sistemas como Bitbucket o GitLab. Después de configurar lo anterior, he añadido un fichero [shippable.yml](https://github.com/pabloalfaro/Car-finder/blob/main/shippable.yml). En este caso, la configuración que he elegido ha sido la de mi primera opcion para Travis, ejecutando los test con la orden de mi fichero Makefile. A diferencia de la anterior, esta vez compruebo más versiones de mi lenguaje.

~~~
language: go

go:
- 1.15
- 1.14
- 1.13
- 1.12
- 1.11
- 1.10

script: 
- make test
~~~

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
- [Tema 3](https://github.com/pabloalfaro/Car-finder/blob/main/Semanas%20anteriores/tema3.md)
- [Tema 4](https://github.com/pabloalfaro/Car-finder/blob/main/Semanas%20anteriores/tema4.md)
