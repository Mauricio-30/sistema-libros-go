package libros

import (
	"sistema-libros-go/domain"
	"sistema-libros-go/funcional"
	"strings"
)

// AgregarLibro retorna un NUEVO slice → función pura (no muta el original).
func AgregarLibro(libros []domain.Libro, nuevo domain.Libro) []domain.Libro {
	resultado := make([]domain.Libro, len(libros), len(libros)+1)
	copy(resultado, libros)
	return append(resultado, nuevo)
}

// EliminarLibro retorna un NUEVO slice sin el libro indicado.
func EliminarLibro(libros []domain.Libro, isbn string) []domain.Libro {
	indice, existe := funcional.BuscarIndice(libros, isbn)
	if !existe {
		return libros
	}
	resultado := make([]domain.Libro, 0, len(libros)-1)
	resultado = append(resultado, libros[:indice]...)
	resultado = append(resultado, libros[indice+1:]...)
	return resultado
}

// BuscarPorAutor usa Filter → función de orden superior.
func BuscarPorAutor(libros []domain.Libro, autor string) []domain.Libro {
	return funcional.FiltrarLibros(libros, func(l domain.Libro) bool {
		return strings.EqualFold(l.Autor, autor)
	})
}

// BuscarPorGenero usa Filter → función de orden superior.
func BuscarPorGenero(libros []domain.Libro, genero string) []domain.Libro {
	return funcional.FiltrarLibros(libros, func(l domain.Libro) bool {
		return strings.EqualFold(l.Genero, genero)
	})
}

// AgruparPorGenero retorna un MAP → demuestra uso de maps.
func AgruparPorGenero(libros []domain.Libro) map[string][]domain.Libro {
	grupos := make(map[string][]domain.Libro)
	for _, l := range libros {
		grupos[l.Genero] = append(grupos[l.Genero], l)
	}
	return grupos
}

// GeneradorISBN crea un closure → demuestra closures.
func GeneradorISBN() func() string {
	contador := 0
	return func() string {
		contador++
		return "ISBN-" + strings.Repeat("0", 3-len(string(rune(contador)))) + string(rune('0'+contador))
	}
}
