package ebook

import (
	"errors"
	"fmt"
)

// GestorLibros administra la colección de libros electrónicos
type GestorLibros struct {
	libros []Libro
}

// NuevoGestorLibros crea una nueva instancia del gestor
func NuevoGestorLibros() *GestorLibros {
	return &GestorLibros{
		libros: make([]Libro, 0),
	}
}

// AgregarLibro añade un libro a la colección
func (g *GestorLibros) AgregarLibro(libro Libro) {
	g.libros = append(g.libros, libro)
	fmt.Println("✅ Libro agregado correctamente.")
}

// BuscarLibroPorTitulo busca un libro por su título
func (g *GestorLibros) BuscarLibroPorTitulo(titulo string) (*Libro, error) {
	for i := range g.libros {
		if g.libros[i].GetTitulo() == titulo {
			return &g.libros[i], nil
		}
	}
	return nil, errors.New("❌ libro no encontrado con ese título")
}

// EliminarLibroPorTitulo elimina un libro por su título
func (g *GestorLibros) EliminarLibroPorTitulo(titulo string) error {
	for i, libro := range g.libros {
		if libro.GetTitulo() == titulo {
			g.libros = append(g.libros[:i], g.libros[i+1:]...)
			fmt.Println("🗑️ Libro eliminado correctamente.")
			return nil
		}
	}
	return errors.New("❌ no se encontró el libro para eliminar")
}

// ListarLibros muestra todos los libros en la colección
func (g *GestorLibros) ListarLibros() {
	if len(g.libros) == 0 {
		fmt.Println("📭 No hay libros registrados.")
		return
	}
	fmt.Println("📚 Lista de libros registrados:")
	for _, libro := range g.libros {
		fmt.Printf("- %s (%s) [%s]\n", libro.GetTitulo(), libro.GetAutor(), libro.GetCategoria())
	}
}
