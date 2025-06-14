package ebook

import (
	"errors"
	
)

// Libro representa un libro electrónico con sus metadatos básicos.
type Libro struct {
	titulo     string
	autor      string
	
	categoria  string
	rutaArchivo string
}

// Constructor para crear un nuevo libro validado
func NuevoLibro(titulo, autor, categoria, rutaArchivo string) (*Libro, error) {
	if titulo == "" || autor == "" {
		return nil, errors.New("todos los campos obligatorios deben estar llenos")
	}
	
	return &Libro{
		titulo:      titulo,
		autor:       autor,
		//
		categoria:   categoria,
		rutaArchivo: rutaArchivo,
	}, nil
}

// ======================= MÉTODOS GET =======================

func (l *Libro) GetTitulo() string {
	return l.titulo
}

func (l *Libro) GetAutor() string {
	return l.autor
}



func (l *Libro) GetCategoria() string {
	return l.categoria
}

func (l *Libro) GetRutaArchivo() string {
	return l.rutaArchivo
}

// ======================= MÉTODOS SET =======================

func (l *Libro) SetTitulo(titulo string) {
	l.titulo = titulo
}

func (l *Libro) SetAutor(autor string) {
	l.autor = autor
}

//func (l *Libro) SetISBN(isbn string) error {
	//if !validarISBN(isbn) {
		//return errors.New("ISBN inválido")
	//}
	//l.isbn = isbn
	//return nil
//}

func (l *Libro) SetCategoria(categoria string) {
	l.categoria = categoria
}

func (l *Libro) SetRutaArchivo(ruta string) {
	l.rutaArchivo = ruta
}


// ======================= VALIDACIONES =======================
// validarISBN verifica si el ISBN tiene un formato válido (simplificado)
//func validarISBN(isbn string) bool {
	//re := regexp.MustCompile(`^\d{10}(\d{3})?$`)
	//return re.MatchString(isbn)
