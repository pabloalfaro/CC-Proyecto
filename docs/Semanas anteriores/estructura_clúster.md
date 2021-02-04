## Estructura del clúster

Hasta ahora la aplicación se componía, de manera resumida, de una clase controladora con las funcionalidades de la aplicación y un api que permitía hacer una serie
de peticiones sobre la aplicación. La manera de trabajar con los datos era con almacenamiento local y sin persistencia de datos. En esta etapa he decido cambiar eso 
y para ello he añadido una base de datos con persistencia de datos para guardar información de los clientes, los coches y los anuncios. La estructura de las tablas se
se encuentra en el archivo [setup.sql](https://github.com/pabloalfaro/Car-finder/blob/main/.docker/setup.sql) y es la siguiente:

~~~
CREATE TABLE if NOT EXISTS usuarios(
username VARCHAR(30) NOT NULL,
nombre VARCHAR(30) NOT NULL,
apellido VARCHAR(30) NOT NULL,
correo VARCHAR(30) NOT NULL,
ciudad VARCHAR(30) NOT NULL,
PRIMARY KEY(username)
)ENGINE=INNODB;

CREATE TABLE if NOT EXISTS coches(
id INT AUTO_INCREMENT,
marca VARCHAR(30) NOT NULL,
modelo VARCHAR(30) NOT NULL,
serie VARCHAR(30) NOT NULL,
potencia INT NOT NULL,
PRIMARY KEY(id)
)ENGINE=INNODB;

CREATE TABLE if NOT EXISTS anuncios(
id INT AUTO_INCREMENT,
usuario VARCHAR(30) NOT NULL,
precio FLOAT NOT NULL,
coche INT NOT NULL,
km INT NOT NULL,
estado VARCHAR(30) NOT NULL,
ciudad VARCHAR(30) NOT NULL,
descripcion VARCHAR(100) NOT NULL,
color VARCHAR(30) NOT NULL,
PRIMARY KEY(id),
FOREIGN KEY(usuario) REFERENCES usuarios(username),
FOREIGN KEY(coche) REFERENCES coches(id)
)ENGINE=INNODB;
~~~

Tanto la base de datos cómo mi aplicación, están implementadas en contenedores. De tal manera que con la configuración de mi docker-compose se pueden levantar
de manera simultánea y pueden interactuar entre si para llevar a cabo una serie de funcionalidades.

