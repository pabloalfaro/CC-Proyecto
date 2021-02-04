# Configuración de los contenedores

Para mi aplicación he configurado dos contenedores: uno con mis funcionalidades y mi api y otro con una base de datos. Para mi api he utilizado un dokerfile
llamado [Dockerfile.despliegue](https://github.com/pabloalfaro/Car-finder/blob/main/Dockerfile.despliegue). Para escribir este fichero he utilizado la información 
de este [post](https://levelup.gitconnected.com/complete-guide-to-create-docker-container-for-your-golang-application-80f3fb59a15e).

## API

Para este contenedor utilizo una [imagen](https://hub.docker.com/_/golang) de alpine para go. Utilizo el directorio de trabajo "build" del contenedor, 
hago un build de la aplicacion, pongo el binario en el directorio "dist" y lo ejecuto. El puerto que exporto es el 8080.

~~~
FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o api ./src/api/api.go

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/api .

# Export necessary port
EXPOSE 8080

# Command to run when starting the container
CMD ["/dist/api"]
~~~

## Base de datos

Para la elección de la base de datos he comparado varias:

### MySQL
Ventajas:
- Facilidad de uso y gran rendimiento
- Facilidad para instalar y configurar
- Soporte multiplataforma
- Soporte SSL

La principal desventaja es la escalabilidad, es decir, no trabaja de manera eficiente con bases de datos muy grandes que superan un determinado tamaño.

### MariaDB
Ventajas:
- Aumento de motores de almacenamiento
- Gran escalabilidad
- Seguridad y rapidez en transacciones
- Extensiones y nuevas características relacionadas con su aplicación para Bases de datos NoSQL.

Las desventajas son algunas incompatibilidades en la migración de MariaDB y MySQL o pequeños atrasos en la liberación de versiones estables.

### Postgresql
Ventajas:
- Control de Concurrencias multiversión (MVCC)
- Flexibilidad en cuanto a lenguajes de programación
- Multiplataforma
- Dispone de una herramienta pgAdmin muy fácil e intuitiva para la administración de las bases de datos.
- Robustez, Eficiencia y Estabilidad.

La principal desventaja es la lentitud para la administración de bases de datos pequeñas ya que está optimizado para gestionar grandes volúmenes de datos.

### Decisión
Viendo estas bases de datos he optado por usar MySQL porque tiene un uso sencillo y para trabajar con bases de datos pequeñas me podía servir. Para un futuro si quisiera 
trabajar con grandes volúmenes de datos podría ser problemático y podría ser necesario cambiar de gestor. En mi caso, en un furturo cercano no preveo trabajar con esas 
cantidades de datos.

