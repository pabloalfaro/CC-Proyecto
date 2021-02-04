# Fichero de composición

El fichero que he usado ha sido este [docker-compose](https://github.com/pabloalfaro/Car-finder/blob/main/docker-compose.yml). Expongo el puerto 8080 del docker en el puerto 8080 
para la api. Para la base de datos hago lo mismo en el puerto 3306. Utilizo la imagen de mi [Dockerfile.despliegue](https://github.com/pabloalfaro/Car-finder/blob/main/Dockerfile.despliegue)
para montar la api y [esta](https://hub.docker.com/_/mysql) de docker para MySQL. La base de datos la monto con el fichero [setup.sql](https://github.com/pabloalfaro/Car-finder/blob/main/.docker/setup.sql)
que se encuentra en el directorio ./docker de mi aplicación. Establezco la contraseña para el usuario root y creo un nuevo usuario. Además pongo como predeterminada mi base de
datos que he creado previamente.

~~~
version: '3.3'
services:
  app:
    build:
      context: .
      dockerfile: Dockerfile.despliegue
    volumes:
      - ./src/:/usr/src/app
    working_dir: /usr/src/app
    ports:
      - "8080:8080"

  db:
    image: mysql
    container_name: godockerDB
    volumes:
      - .docker/setup.sql:/docker-entrypoint-initdb.d/setup.sql
    environment:
      MYSQL_USER: docker
      MYSQL_ROOT_PASSWORD: password
      MYSQL_PASSWORD: password
      MYSQL_DATABASE: carfinder

    ports:
      - "3306:3306"
    restart: always
~~~
