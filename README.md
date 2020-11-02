# Car finder (En desarrollo)
Soy Pablo Alfaro, alumno de la asignatura de Cloud Computing del Máster en Ingeniería Informática de la Universidad de Granada. En este repositorio se encuentra el proyecto de la asignatura.


## Definición de arquitectura y lenguaje
Para mi proyecto he decidido usar Go porque he visto que se utiliza mucho para el desarrollo web backend. Además, he visto que cubría mis necesidades, como por ejemplo la de tener que usar variables privadas.

Para decidir la arquitectura tengo antes que analizar la aplicación. En mi caso veo que voy a tener una serie de elementos diferenciados: la gestión de los usuarios, la de los anuncios, la de los distintos modelos de coches y por último la mensajería entre usuarios. Estos son los que componen la idea inicial del proyecto pero no descarto nuevas funcionalidades. Además, los elementos mencionados almacenarán datos diferentes, relacionados con las características que gestionen. Estos bloques tendrán que interactuar entre sí.

Teniendo en cuenta lo anterior, me he decantado por una arquitectura de microservicios. De esta manera podría tener un servicio encargado de gestionar cada uno de los 4 bloques mencionados anteriormente. Estos servicios tendían su BD correspondiente en la que almacenarían los datos que gestionan. La ventaja que tiene trabajar con esta arquitectura en mi proyecto es la modularidad. Gracias a esto, si en un futuro necesito añadir nuevas funcionalidades me será sencillo. Los diferentes bloques necesitan interactuar entre sí, por lo que los servicios se tendrán que comunicar, en mi caso, de manera asíncrona.

## Historias de usuario

En mi caso he pensado en 3 posibles usuarios, estos son:

- El cliente que ha encargado el proyecto: [Cliente](https://github.com/pabloalfaro/Car-finder/issues?q=is%3Aissue+is%3Aopen+label%3Acliente)
- El comprador de los coches: [Comprador](https://github.com/pabloalfaro/Car-finder/issues?q=is%3Aissue+is%3Aopen+label%3Acomprador)
- El vendedor de los coches: [Vendedor](https://github.com/pabloalfaro/Car-finder/issues?q=is%3Aissue+is%3Aopen+label%3Avendedor)


## Documentación adicional
- [Roadmap del proyecto](https://github.com/pabloalfaro/Car-finder/blob/main/roadmap.md)
- [Kanban](https://github.com/pabloalfaro/Car-finder/projects/1)
- [Clases programadas](https://github.com/pabloalfaro/Car-finder/tree/main/src/main)
- [Clases compiladas](https://github.com/pabloalfaro/Car-finder/blob/main/Documentaci%C3%B3n%20adicional/verificaci%C3%B3n%20del%20c%C3%B3digo.png)
- [Producto mínimamente viable 1](https://github.com/pabloalfaro/Car-finder/milestone/3)
- [Issues relacionados con el Hito1 de la asignatura](https://github.com/pabloalfaro/Car-finder/milestone/2)


## Contenido anterior
Aquí dejo los links a los contenidos de las semanas anteriores:

- [Semana1](https://github.com/pabloalfaro/Car-finder/blob/main/Semanas%20anteriores/tema1.md)
