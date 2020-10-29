# Car finder
Soy Pablo Alfaro, alumno de la asignatura de Cloud Computing del Máster en Ingeniería Informática de la Universidad de Granada. En este repositorio se encuentra el proyecto de la asignatura.


## Definición de arquitectura y lenguaje
Para mi proyecto he decidido usar Go porque he visto que se utiliza mucho para el desarrollo web backend. Además, he visto que cubría mis necesidades, como por ejemplo la de tener que usar variables privadas.

En cuanto a la arquitectura, me he decantado por una de microservicios. La ventaja que tiene trabajar con esta arquitectura y la que han hecho decidirme por ella ha sido la modularidad. Esto quiere decir que los servicios se pueden implementar de manera independiente y se pueden desplegar de la misma manera. En mi caso esta arquitectura se puede aplicar con un servicio que gestione los anuncios de los coches, otro que gestione los usuarios, otro para gestionar el listado de coches y por último uno que se encargue de la mensajería entre usuarios. Cada servicio tendría su base de datos correspondiente.

## Historias de usuario

En mi caso he pensado en 4 posibles usuarios, estos son:

- El cliente que ha encargado el proyecto: [Cliente](https://github.com/pabloalfaro/Car-finder/issues?q=is%3Aissue+is%3Aopen+label%3Acliente)
- El comprador de los coches: [Comprador](https://github.com/pabloalfaro/Car-finder/issues?q=is%3Aissue+is%3Aopen+label%3Acomprador)
- El vendedor de los coches: [Vendedor](https://github.com/pabloalfaro/Car-finder/issues?q=is%3Aissue+is%3Aopen+label%3Avendedor)


## Documentación adicional
- [Roadmap del proyecto](https://github.com/pabloalfaro/Car-finder/blob/main/roadmap.md)
- [Kanban](https://github.com/pabloalfaro/Car-finder/projects/1)
- [Clases programadas](https://github.com/pabloalfaro/Car-finder/tree/main/src/main)
- [Producto mínimamente viable 1] (https://github.com/pabloalfaro/Car-finder/milestone/3)

- [Issues relacionados con el Hito1 de la asignatira] (https://github.com/pabloalfaro/Car-finder/milestone/2)

Las clases que he programado han sido una definición inicial de los objetos con los que voy a trabajar. Además, he definido los métodos get y set. Algunos de estos métodos tendré que eliminarlos porque hacen referencia a atributos inmutables. Hago referencia a esto en el [issue #19](https://github.com/pabloalfaro/Car-finder/issues/19).

## Contenido anterior
Aquí dejo los links a los contenidos de las semanas anteriores:

- [Semana1](https://github.com/pabloalfaro/Car-finder/blob/main/Semanas%20anteriores/tema1.md)
