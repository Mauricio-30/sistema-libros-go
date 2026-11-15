package funcional

import "sistema-libros-go/domain"

// FiltrarLibros recibe un predicado (función) → función de orden superior.
func FiltrarLibros(libros []domain.Libro, predicado func(domain.Libro) bool) []domain.Libro {
	var resultado []domain.Libro
	for _, l := range libros {
		if predicado(l) {
			resultado = append(resultado, l)
		}
	}
	return resultado
}

// MapearTitulos transforma cada libro → función de orden superior.
func MapearTitulos(libros []domain.Libro, transformar func(domain.Libro) string) []string {
	resultado := make([]string, len(libros))
	for i, l := range libros {
		resultado[i] = transformar(l)
	}
	return resultado
}

// BuscarIndice retorna el índice y un booleano → RETORNO MÚLTIPLE.
func BuscarIndice(libros []domain.Libro, isbn string) (int, bool) {
	for i, l := range libros {
		if l.ISBN == isbn {
			return i, true
		}
	}
	return -1, false
}
