help:
	echo "Acciones válidas: install build utest"

install:
	go install github.com/pabloalfaro/Car-finder/src/anuncio
	go install github.com/pabloalfaro/Car-finder/src/coche
	go install github.com/pabloalfaro/Car-finder/src/controlador
	go install github.com/pabloalfaro/Car-finder/src/mensaje
	go install github.com/pabloalfaro/Car-finder/src/usuario
	
build:
	go build ./src/controlador
	go build ./src/anuncio
	go build ./src/coche
	go build ./src/mensaje
	go build ./src/usuario

utest: 
	go test ./test
