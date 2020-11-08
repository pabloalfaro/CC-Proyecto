# Car finder (En desarrollo)
Soy Pablo Alfaro, alumno de la asignatura de Cloud Computing del Máster en Ingeniería Informática de la Universidad de Granada. En este repositorio se encuentra el proyecto de la asignatura.


## Gestor de tareas

Para mi proyecto he utilizado [Makefile](https://github.com/pabloalfaro/Car-finder/blob/main/Makefile). Me he decido por este gestor porque es uno de los más utilizados. Cualquiera que se encuentre con este tipo de ficheros sabe como uitilizar las instrucciones básicas sin tener que buscar información al respecto.

## Biblioteca de aserciones

Mi proyecto está escrito en Go, que tiene su propia biblioteca de aserciones `testing`. He utilizado esta biblioteca porque es la primera vez que utilizo este lenguaje de programación y quiero intentar usar las herramientas propias de Go para estos casos. Además, he visto que existen otras bibliotecas pero de manera estandar se utiliza `testintg`.

## Ejecución de los tests

En este apartado ocurre algo parecido al anterior. Go tiene una manera propia de ejecutar los tests. Utilizando el comando `go test` ejecutas los test que tengas en el directorio dónde lo estés ejecutando. Otra manera de ejecutrarlos sería con `go test ./ruta_de_tests`, de esta manera puedes elegir el directorio dónde están los test que quieras ejecutar. Los archivos de Go de test se identifican porque tienen este tipo de nombres: `nombre_test.go`. Estos archivos puedes guardarlos junto con la clase que testean o los puedes guardar en otro directorio a parte. Go no permite normalmente que tengas ficheros de dos `package` distintos, pero con los tests hace una excepción.

## Historias de usuario

En mi caso he pensado en 3 posibles usuarios, estos son:

- El cliente que ha encargado el proyecto: [Cliente](https://github.com/pabloalfaro/Car-finder/issues?q=is%3Aissue+is%3Aopen+label%3Acliente)
- El comprador de los coches: [Comprador](https://github.com/pabloalfaro/Car-finder/issues?q=is%3Aissue+is%3Aopen+label%3Acomprador)
- El vendedor de los coches: [Vendedor](https://github.com/pabloalfaro/Car-finder/issues?q=is%3Aissue+is%3Aopen+label%3Avendedor)



## Documentación adicional
- [Roadmap del proyecto](https://github.com/pabloalfaro/Car-finder/blob/main/roadmap.md)
- [Kanban](https://github.com/pabloalfaro/Car-finder/projects/1)
- [Clases programadas](https://github.com/pabloalfaro/Car-finder/tree/main/src)
- [Tests unitarios](https://github.com/pabloalfaro/Car-finder/tree/main/test)
- [PMV 1](https://github.com/pabloalfaro/Car-finder/milestone/3)
- [PMV 2](https://github.com/pabloalfaro/Car-finder/milestone/4)
- [Issues relacionados con los Objetivos 1 de la asignatura](https://github.com/pabloalfaro/Car-finder/milestone/2)


## Contenido anterior
Aquí dejo los links a los contenidos de las semanas anteriores:

- [Tema 1](https://github.com/pabloalfaro/Car-finder/blob/main/Semanas%20anteriores/tema1.md)
- [Tema 2](https://github.com/pabloalfaro/Car-finder/blob/main/Semanas%20anteriores/tema2.md)
