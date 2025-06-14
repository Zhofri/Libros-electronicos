package user

import (
	"errors"
	"fmt"
)

// GestorUsuarios administra la lista de usuarios del sistema
type GestorUsuarios struct {
	usuarios []Usuario
}

// NuevoGestorUsuarios crea un nuevo gestor de usuarios
func NuevoGestorUsuarios() *GestorUsuarios {
	return &GestorUsuarios{
		usuarios: make([]Usuario, 0),
	}
}

// RegistrarUsuario agrega un nuevo usuario al sistema si no existe previamente
func (g *GestorUsuarios) RegistrarUsuario(u Usuario) error {
	for _, usuario := range g.usuarios {
		if usuario.GetCorreo() == u.GetCorreo() {
			return errors.New("❌ ya existe un usuario con ese correo")
		}
	}
	g.usuarios = append(g.usuarios, u)
	fmt.Println("✅ Usuario registrado correctamente.")
	return nil
}

// AutenticarUsuario verifica si el correo y contraseña coinciden con un usuario registrado
func (g *GestorUsuarios) AutenticarUsuario(correo, contrasena string) (*Usuario, error) {
	for i := range g.usuarios {
		if g.usuarios[i].GetCorreo() == correo && g.usuarios[i].contrasena == contrasena {
			return &g.usuarios[i], nil
		}
	}
	return nil, errors.New("❌ usuario o contraseña incorrectos")
}

// ListarUsuarios imprime los usuarios registrados (solo correos y roles)
func (g *GestorUsuarios) ListarUsuarios() {
	if len(g.usuarios) == 0 {
		fmt.Println("📭 No hay usuarios registrados.")
		return
	}
	fmt.Println("👥 Usuarios registrados:")
	for _, u := range g.usuarios {
		fmt.Printf("- %s (%s)\n", u.GetCorreo(), u.GetNombre())
	}
}

// ObtenerUsuarios devuelve todos los usuarios (para almacenamiento)
func (g *GestorUsuarios) ObtenerUsuarios() []Usuario {
	return g.usuarios
}

// CargarUsuarios reemplaza la lista actual por una lista cargada desde archivo
func (g *GestorUsuarios) CargarUsuarios(usuarios []Usuario) {
	g.usuarios = usuarios
}
