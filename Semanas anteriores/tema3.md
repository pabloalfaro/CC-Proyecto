## Gestor de tareas

Para mi proyecto he utilizado [Makefile](https://github.com/pabloalfaro/Car-finder/blob/main/Makefile). Me he decido por este gestor porque es uno de los más utilizados. Cualquiera que se encuentre con este tipo de ficheros sabe como uitilizar las instrucciones básicas sin tener que buscar información al respecto. He implementado una funcionalidad de test, una funcionalidad "fmt" que hace que el código programado sea más legible, la de install y un all. Para escribir mi fichero he seguido algunas buenas prácticas para Golang descritas en esta [Web](https://ops.tips/blog/minimal-golang-makefile/).

## Biblioteca de aserciones 

La biblioteca de aserciones que he utilizado ha sido [Gomega](https://github.com/onsi/gomega). En este caso la elección ha venido dada por ser el recomendado por mi marco de pruebas. En mi caso he utilizado aserciones para comprobar las salidas de mis métodos utilizando la funcion `Expect()`, `To()`, `BeTrue()` y `BeFalse()`. Estos métodos devuelven un valor de True cuando se ha completado con éxito y un valor de False en el caso contrario.

## Marco de pruebas 

Me he decidido por tests de tipo bdd, este tipo de test se han de implementar para comprobar que se cumple con las funcionalidades descritas en las distintas HU. En este caso se hace de una manera casi literal, intentando que los tests se parezcan lo máximo posible al lenguaje natural. He hecho pruebas con [Godog](https://github.com/cucumber/godog) y con [Ginkgo](https://github.com/onsi/ginkgo). He probado ambos frameworks, pero con el de Godog no conseguí correr los tests, las implementaciones de pueden ver en el commit [#47](https://github.com/pabloalfaro/Car-finder/commit/afea01ddabb865584b1fa0807eed2b5aacf6453f). Después de intentarlo con Godog probé con Ginko, que es con el que finalmente me he quedado. Tras haber usado los dos y aunque hubiese podido correr los de Godog, me hubiese quedado con Ginkgo. Me ha resultado menos laborioso implementar los test y personalmente veo que el código es más claro en este caso. 

## Ejecución de los tests 

Ginkgo tiene una manera pripia de ejecutar los tests pero se puede hacer con la herramienta de Go, utilizando el comando `go test`. De esta manera ejecutas los test que tengas en el directorio dónde lo estés ejecutando. Otra método de ejecutrarlos sería con `go test ./ruta_de_tests`, de esta forma puedes elegir el directorio dónde están los test que quieras ejecutar. Los archivos de Go de test se identifican porque tienen este tipo de nombres: `nombre_test.go`. Estos archivos puedes guardarlos junto con la clase que testean o los puedes guardar en otro directorio a parte. Go no permite normalmente que tengas ficheros de dos `package` distintos, pero con los tests hace una excepción. En mi caso he guardado los test junto con la clase a la que hacen referencia.
