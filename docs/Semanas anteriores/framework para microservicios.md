## Comparativa entre framework para microservicios en GO

Las herramientas que he estado investigando para el desempeño de esta tarea han sido las siguientes:
- Gin-gonic
- Macaron
-	Echo
-	Iris
-	Beego


### [Gin-gonic](https://github.com/gin-gonic/gin):
Es un framework web escrito en golang. Cuenta con una API similar a la de martini con un rendimiento mucho mejor, hasta 40 veces más rápido.

-	Extremadamente simple de usar.
-	Diseño no intrusivo.
-	Buena integración con otros paquetes Golang.
-	Enrutamiento impresionante.
-	Diseño modular - Fácil de añadir y remover funcionalidades.
-	Muy buen uso de handlers/middlewares.
-	Sirviendo documentos por defecto (e.g. para servir aplicaciones AngularJS en modo HTML5).

He encontrado información adicional en las siguientes páginas:
- https://gin-gonic.com/es/docs/
- https://www.cibernomadas.es/desarrollo/go-con-gin-el-mega-tutorial/


### [Macaron](https://github.com/go-macaron/macaron):
-	Tiene un fuerte enrutamiento con suburl.
-	Es flexible en la combinación de rutas.
-	Enrutadores de grupo anidados ilimitados. 
-	Integra directamente los servicios existentes.
-	Puede cambiar dinámicamente las plantillas en tiempo de ejecución.
-	Permite usar plantillas en memoria y archivos estáticos. 
-	Las funcionalidades son fáciles de conectar y desconectar gracias al diseño modular.
-	Permite la inyección de dependencias usando [inject](https://github.com/codegangsta/inject). 

He encontrado información adicional en la siguiente página:
- https://go-macaron.com/

### [Echo](https://github.com/labstack/echo):
-	Tiene un enrutador HTTP optimizado que prioriza las rutas de manera inteligente.
-	Permite crear API RESTful robustas y escalables.
-	Premite agupar APIs.
-	Marco de middleware extensible.
-	Permite definir middleware a nivel de raíz, grupo o ruta.
-	Enlace de datos para JSON, XML.
-	Contiene funciones útiles para enviar una variedad de respuestas HTTP.
-	Tiene un manejo centralizado de errores HTTP.
-	Define el formato para el logger.
-	Tiene soporte para HTTP/2.

He encontrado información adicional en la siguiente página:
- https://medium.com/cuddle-ai/building-microservice-using-golang-echo-framework-ff10ba06d508


### [Iris](https://github.com/kataras/iris):
Es un framework web rápido y simple, pero con muchas funcionalidades y muy eficiente para Go. Proporciona una base fácil de usar para crear un sitio web o una API.

-	API similar a Sinatra. Uso familiar si se ha trabajado con expressjs.
-	Iris funciona más rápido que cualquier otro framework HTTP incluyendo Gin.
-	Iris se basa en el paquete http estándar. Es totalmente compatible con todos los middlewares existentes.
-	Iris admite parámetros de ruta dinámica y funciones en línea.
-	Permite ejecutar tantos servidores web Iris como como quieras en uno o más Redis Clusters.
-	Iris Context admite todos los formatos principales: texto, HTML, Markdown, XML, YAML, binario, JSON, JSONP, Problem, Protocol Buffers, Message Pack.
-	Iris proporciona un adaptador para un increíble conjunto de pruebas de terceros diseñado para facilitar la lectura.

He encontrado información adicional en las siguientes páginas:
- https://www.iris-go.com/
- https://godoc.org/gopkg.in/iris-framework/iris.v6
- https://docs.iris-go.com/iris/contents/testing


### [Beego](https://github.com/beego/beego/):
Beego se utiliza para el desarrollo rápido de aplicaciones empresariales en Go, incluidas las API RESTful, las aplicaciones web y los servicios de backend. Está inspirado en Tornado, Sinatra y Flask. beego tiene algunas características específicas de Go, como interfaces y estructuras embebidas.

Tiene soporte RESTful, modelo MVC y usa la herramienta bee para crear aplicaciones que incluyen compilación en caliente de código, pruebas automatizadas y empaquetado e implementación automatizada.

Permite el enrutamiento y monitoreo inteligente. Puede monitorear su QPS, el uso de memoria y CPU, y el estado de la rutina. Proporciona un control total de las aplicaciones en línea.

Tiene módulos integrados que incluyen control de sesión, almacenamiento en caché, registro, análisis de configuración, supervisión de rendimiento, manejo de contexto, soporte de ORM y simulación de solicitudes.

Contiene un paquete nativo Go http para manejar las solicitudes y la concurrencia de goroutine.

He encontrado información adicional en las siguientes páginas:
- https://beego.me/
- https://bestow.info/get-going-with-beego-framework/

## Framework elegido

Después de ver diferentes alternativas he optado por [Iris](https://github.com/kataras/iris). Viendo que era el más rápido y su familiaridad en el caso de haber usado expressjs creo que era una buena opción. He revisado también su repo y he visto que estaba actualizado, además, tenía una extensa docuemtación de uso que facilita el iniciarse en este framework.
