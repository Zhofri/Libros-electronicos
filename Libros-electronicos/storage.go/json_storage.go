package storage

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/Zhofri/Libros-electronicos/ebook"
)

// Storage define las operaciones para persistencia de libros
type Storage interface {
	Guardar(libros []ebook.Libro) error
	Cargar() ([]ebook.Libro, error)
}

// JSONStorage implementa Storage usando archivos .json
type JSONStorage struct {
	rutaArchivo string
}

// NuevoJSONStorage crea una instancia de JSONStorage
func NuevoJSONStorage(ruta string) *JSONStorage {
	return &JSONStorage{rutaArchivo: ruta}
}

// Guardar guarda los libros en un archivo JSON
func (js *JSONStorage) Guardar(libros []ebook.Libro) error {
	archivo, err := os.Create(js.rutaArchivo)
	if err != nil {
		return errors.New("no se pudo crear el archivo de almacenamiento")
	}
	defer archivo.Close()

	encoder := json.NewEncoder(archivo)
	encoder.SetIndent("", "  ") // para que sea más legible

	if err := encoder.Encode(libros); err != nil {
		return errors.New("error al codificar los datos en JSON")
	}
	return nil
}

// Cargar lee los libros desde un archivo JSON
func (js *JSONStorage) Cargar() ([]ebook.Libro, error) {
	var libros []ebook.Libro

	archivo, err := os.Open(js.rutaArchivo)
	if err != nil {
		// Si el archivo no existe, retornamos una lista vacía
		if os.IsNotExist(err) {
			return libros, nil
		}
		return nil, errors.New("no se pudo abrir el archivo de almacenamiento")
	}
	defer archivo.Close()

	decoder := json.NewDecoder(archivo)
	if err := decoder.Decode(&libros); err != nil {
		return nil, errors.New("error al decodificar los datos JSON")
	}
	return libros, nil
}
