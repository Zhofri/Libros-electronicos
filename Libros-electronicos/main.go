package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Zhofri/Libros-electronicos/ebook"
	"github.com/Zhofri/Libros-electronicos/user"
)

var lector = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("📚 Bienvenido al Sistema de Gestión de Libros Electrónicos 📚")

	gestorUsuarios := user.NuevoGestorUsuarios()
	gestorLibros := ebook.NuevoGestorLibros()

	for {
		fmt.Println("\nSelecciona una opción:")
		fmt.Println("1. Registrar usuario")
		fmt.Println("2. Iniciar sesión")
		fmt.Println("3. Salir")

		opcion := leerEntrada("Opción")

		switch opcion {
		case "1":
			registrarUsuario(gestorUsuarios)
		case "2":
			usuario, err := login(gestorUsuarios)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("\n🎉 ¡Bienvenido %s!\n", usuario.GetNombre())
			menuUsuario(usuario, gestorLibros)
		case "3":
			fmt.Println("👋 Hasta luego.")
			return
		default:
			fmt.Println("❌ Opción no válida.")
		}
	}
}

func registrarUsuario(gestor *user.GestorUsuarios) {
	nombre := leerEntrada("Nombre")
	correo := leerEntrada("Correo")
	contra := leerEntrada("Contraseña")
	//rol := leerEntrada("Rol (admin/usuario)")

	u, err := user.NuevoUsuario(nombre, correo, contra, "user") // Asignamos rol "usuario" por defecto
	// Si se desea implementar roles, se puede agregar lógica para asignar "admin" o "usuario"		
	if err != nil {
		fmt.Println("❌ Error:", err)
		return
	}

	err = gestor.RegistrarUsuario(*u)
	if err != nil {
		fmt.Println("❌ Error:", err)
	}
}

func login(gestor *user.GestorUsuarios) (*user.Usuario, error) {
	correo := leerEntrada("Correo")
	contra := leerEntrada("Contraseña")

	u, err := gestor.AutenticarUsuario(correo, contra)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func menuUsuario(u *user.Usuario, gestorLibros *ebook.GestorLibros) {
	for {
		fmt.Println("\n--- Menú de Usuario ---")
		fmt.Println("1. Agregar libro")
		fmt.Println("2. Listar libros")
		fmt.Println("3. Buscar libro por título")
		fmt.Println("4. Eliminar libro por título")
		fmt.Println("5. Cerrar sesión")

		opcion := leerEntrada("Opción")

		switch opcion {
		case "1":
			titulo := leerEntrada("Título")
			autor := leerEntrada("Autor")
			//isbn := leerEntrada("ISBN")
			categoria := leerEntrada("Categoría")
			ruta := leerEntrada("Ruta del archivo")

			
			libro, err := ebook.NuevoLibro(titulo, autor, categoria, ruta)
			if err != nil {
				fmt.Println("❌ Error:", err)
				continue
			}
			gestorLibros.AgregarLibro(*libro)

		case "2":
			gestorLibros.ListarLibros()

		case "3":
			titulo := leerEntrada("Título a buscar")
			libro, err := gestorLibros.BuscarLibroPorTitulo(titulo)
			if err != nil {
				fmt.Println("❌", err)
			} else {
				fmt.Printf("🔍 Libro encontrado: %s (%s) - %s\n", libro.GetTitulo(), libro.GetAutor(), libro.GetCategoria())
			}

		case "4":
			titulo := leerEntrada("Título del libro a eliminar")
			err := gestorLibros.EliminarLibroPorTitulo(titulo)
			if err != nil {
				fmt.Println("❌", err)
			}

		case "5":
			fmt.Println("🔒 Cerrando sesión...")
			return

		default:
			fmt.Println("❌ Opción no válida.")
		}
	}
}

func leerEntrada(campo string) string {
	fmt.Printf("Ingrese %s: ", campo)
	texto, _ := lector.ReadString('\n')
	return strings.TrimSpace(texto)
}
