package usuarios

import (
	"sistema-libros-go/domain"
	"strings"
)

// AgregarUsuario retorna un NUEVO slice → función pura.
func AgregarUsuario(usuarios []domain.Usuario, nuevo domain.Usuario) []domain.Usuario {
	resultado := make([]domain.Usuario, len(usuarios), len(usuarios)+1)
	copy(resultado, usuarios)
	return append(resultado, nuevo)
}

// BuscarPorEmail retorna (usuario, bool) → RETORNO MÚLTIPLE.
func BuscarPorEmail(usuarios []domain.Usuario, email string) (domain.Usuario, bool) {
	for _, u := range usuarios {
		if strings.EqualFold(u.Email, email) {
			return u, true
		}
	}
	return domain.Usuario{}, false
}

// ListarNombres retorna un NUEVO slice de strings.
func ListarNombres(usuarios []domain.Usuario) []string {
	resultado := make([]string, len(usuarios))
	for i, u := range usuarios {
		resultado[i] = u.Nombre
	}
	return resultado
}
