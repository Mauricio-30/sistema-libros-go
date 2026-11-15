package domain

// Libro representa un libro electrónico.
type Libro struct {
	ISBN   string
	Titulo string
	Autor  string
	Genero string
	Anio   int
}

// Usuario representa a un usuario del sistema.
type Usuario struct {
	ID     string
	Nombre string
	Email  string
}