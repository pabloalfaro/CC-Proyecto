package anuncio

type Anuncio struct {
	id					int
	usuario     string
	precio      float32
	coche       int
	km          int
	estado      string
	ciudad      string
	descripcion string
	color       string
}

//Constructor anuncio
func NewAnuncio(id int, u string, p float32, coc int, k int, e string, ciu string, d string, col string) Anuncio {
	var anuncio Anuncio

	anuncio.id = id
	anuncio.usuario = u
	anuncio.precio = p
	anuncio.coche = coc
	anuncio.km = k
	anuncio.estado = e
	anuncio.ciudad = ciu
	anuncio.descripcion = d
	anuncio.color = col

	return anuncio
}

func GetId(a Anuncio) int {
	return a.id
}

func GetUsuario(a Anuncio) string {
	return a.usuario
}

//Get y set para Precio
func GetPrecio(a Anuncio) float32 {
	return a.precio
}

func SetPrecio(a Anuncio, p float32) {
	a.precio = p
}

//Get para Coche
func GetCoche(a Anuncio) int {
	return a.coche
}

//Get y set para km
func GetKm(a Anuncio) int {
	return a.km
}

func SetKm(a Anuncio, k int) {
	a.km = k
}

//Get y set para estado
func GetEstado(a Anuncio) string {
	return a.estado
}

func SetEstado(a Anuncio, e string) {
	a.estado = e
}

//Get y set para ciudad
func GetCiudad(a Anuncio) string {
	return a.ciudad
}

func SetCiudad(a Anuncio, c string) {
	a.ciudad = c
}

//Get y set para descripcion
func GetDescripcion(a Anuncio) string {
	return a.descripcion
}

func SetDescripcion(a Anuncio, d string) {
	a.descripcion = d
}

//Get y set para color
func GetColor(a Anuncio) string {
	return a.color
}

func SetColorA(a Anuncio, c string) {
	a.color = c
}
