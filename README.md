# Sistema de Gestión de Libros Electrónicos

Proyecto individual desarrollado en Go (Golang) aplicando el paradigma funcional y los temas vistos en clase.

# Temas aplicados
- **Funciones y Retorno Múltiple:** Uso de `(valor, bool)` para validar entradas en `utils/entrada.go`.
- **Closures:** Generador automático de códigos ISBN en `libros/servicio.go`.
- **Slices:** Uso de listas dinámicas (`[]domain.Libro`) en lugar de arrays, para permitir un registro ilimitado de libros.
- **Maps:** Uso de mapas (`map[string][]domain.Libro`) para agrupar los libros por género en `libros/servicio.go`.
- **Programación Funcional:** Funciones puras (no mutan los datos originales) y funciones de orden superior (`FiltrarLibros`, `MapearTitulos`).

##  Ejecución
Para correr el proyecto, ejecutar en la terminal:
```bash
go run main.go
