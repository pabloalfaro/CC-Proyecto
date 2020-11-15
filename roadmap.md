# Roadmap

Este documento se irá actualizando conforme avance el proyecto. De esta forma, se iran añadiendo las funcionalidades implementadas y los planes futuros. Las funcionalidades se irán agrupando en los diferentes *productos mínimamente viables*.

## [PMV 1](https://github.com/pabloalfaro/Car-finder/milestone/3) (Finalizado)

En primer lugar me planteé cuál era la funcionalidad más básica de la aplicación. En mi caso, es una aplicación de venta de coches por lo que llegué a la conclusión de que lo primero que debía hacer era añadir la funcionalidad de publicar anuncios. Esto forma parte de la siguiente historia de usuario: [HU Crear anuncios](https://github.com/pabloalfaro/Car-finder/issues/8).

La primera actividad relacionada con esa historia era la de crear las clases con las que se iba a trabajar ([ISSUE crear las clases](https://github.com/pabloalfaro/Car-finder/issues/24)). Además, relacionado con lo anterior, añadí funcionalidades a las clases ([ISSUE añadir funcionalidad](https://github.com/pabloalfaro/Car-finder/issues/21)). Al princio creé todos los métodos get y set, pero no todos eran necesarios al haber variables inmutables ([ISSUE corregir métodos](https://github.com/pabloalfaro/Car-finder/issues/19)). 

Teniendo ya todo lo anterior programado, demostré que no había errores de compilacion con esta [imagen](https://github.com/pabloalfaro/Car-finder/blob/main/Documentaci%C3%B3n%20adicional/verificaci%C3%B3n%20del%20c%C3%B3digo.png), esto forma parte del siguiente issue ([ISSUE demostrar clases compiladas](https://github.com/pabloalfaro/Car-finder/issues/22)).

En este punto sólo tenía una captura de las clases compiladas para verificar que el código no tiene errores, esto no es suficiente. Teniendo ya algunas clases creadas, habría que programar los tests para comprobar el buen funcionamiento de las mismas. Esta será la prioridad en esta etapa, una vez que tenga los test continuaré con las funcionalidades relacionadas con [HU Crear anuncios](https://github.com/pabloalfaro/Car-finder/issues/8).


## [PMV 2](https://github.com/pabloalfaro/Car-finder/milestone/4) (En desarrollo)

El objetivo principal será crear un Doker dónde se desplegará la aplicación.

Las próximas funcionalidades que planteo son las relacionadas con las bases de datos. Tendré que crear una BD para almacenar los diferentes usuarios, otra para los modelos de los coches y una para los propios anuncios.

Cuando termine con la base de datos empezaré con las distintas funcionalidades para gestionar anuncios y hacer las búsquedas.

