package main

import (
	"fmt"
	"sistema-libros-go/domain"
	"sistema-libros-go/libros"
	"sistema-libros-go/usuarios"
	"sistema-libros-go/utils"
)

func main() {
	listaLibros := []domain.Libro{}
	listaUsuarios := []domain.Usuario{}

	// Aquí usamos el CLOSURE que creamos en libros/servicio.go
	generarISBN := libros.GeneradorISBN()

	for {
		fmt.Println("\n=== SISTEMA DE GESTIÓN DE LIBROS ===")
		fmt.Println("1. Registrar libro")
		fmt.Println("2. Listar libros")
		fmt.Println("3. Buscar por autor")
		fmt.Println("4. Buscar por género")
		fmt.Println("5. Agrupar por género (usando MAP)")
		fmt.Println("6. Registrar usuario")
		fmt.Println("7. Listar usuarios")
		fmt.Println("8. Eliminar libro")
		fmt.Println("0. Salir")
		fmt.Print("Opción: ")

		opcion, ok := utils.LeerEntero("")
		if !ok {
			fmt.Println("Opción inválida. Intenta de nuevo.")
			continue
		}

		switch opcion {
		case 1:
			titulo, _ := utils.LeerTexto("Título: ")
			autor, _ := utils.LeerTexto("Autor: ")
			genero, _ := utils.LeerTexto("Género: ")
			anio, _ := utils.LeerEntero("Año: ")

			nuevo := domain.Libro{
				ISBN:   generarISBN(), // El closure entra en acción
				Titulo: titulo,
				Autor:  autor,
				Genero: genero,
				Anio:   anio,
			}
			listaLibros = libros.AgregarLibro(listaLibros, nuevo)
			fmt.Println(" Libro registrado con el código:", nuevo.ISBN)

		case 2:
			if len(listaLibros) == 0 {
				fmt.Println(" No hay libros registrados.")
				continue
			}
			fmt.Println("\n--- LISTA DE LIBROS ---")
			for _, l := range listaLibros {
				fmt.Printf("[%s] %s - %s (%s, %d)\n", l.ISBN, l.Titulo, l.Autor, l.Genero, l.Anio)
			}

		case 3:
			autor, _ := utils.LeerTexto("Autor a buscar: ")
			encontrados := libros.BuscarPorAutor(listaLibros, autor)
			if len(encontrados) == 0 {
				fmt.Println(" No se encontraron libros de ese autor.")
			} else {
				for _, l := range encontrados {
					fmt.Printf("- %s (%s)\n", l.Titulo, l.Autor)
				}
			}

		case 4:
			genero, _ := utils.LeerTexto("Género a buscar: ")
			encontrados := libros.BuscarPorGenero(listaLibros, genero)
			if len(encontrados) == 0 {
				fmt.Println(" No se encontraron libros de ese género.")
			} else {
				for _, l := range encontrados {
					fmt.Printf("- %s\n", l.Titulo)
				}
			}

		case 5:
			grupos := libros.AgruparPorGenero(listaLibros)
			if len(grupos) == 0 {
				fmt.Println(" No hay libros para agrupar.")
				continue
			}
			for genero, librosGenero := range grupos {
				fmt.Printf("\nGénero: %s (%d libros)\n", genero, len(librosGenero))
				for _, l := range librosGenero {
					fmt.Printf("  - %s\n", l.Titulo)
				}
			}

		case 6:
			nombre, _ := utils.LeerTexto("Nombre: ")
			email, _ := utils.LeerTexto("Email: ")
			id, _ := utils.LeerTexto("ID: ")
			nuevo := domain.Usuario{ID: id, Nombre: nombre, Email: email}
			listaUsuarios = usuarios.AgregarUsuario(listaUsuarios, nuevo)
			fmt.Println(" Usuario registrado.")

		case 7:
			nombres := usuarios.ListarNombres(listaUsuarios)
			if len(nombres) == 0 {
				fmt.Println(" No hay usuarios registrados.")
				continue
			}
			fmt.Println("\n--- USUARIOS ---")
			for _, n := range nombres {
				fmt.Println("-", n)
			}

		case 8:
			isbn, _ := utils.LeerTexto("ISBN a eliminar: ")
			listaLibros = libros.EliminarLibro(listaLibros, isbn)
			fmt.Println(" Libro eliminado (si existía).")

		case 0:
			fmt.Println("👋 ¡Hasta luego!")
			return

		default:
			fmt.Println("⚠️ Opción no válida.")
		}
	}
}
