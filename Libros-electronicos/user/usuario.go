package user

import (
	"errors"
	"regexp"
)

// Usuario representa un usuario del sistema
type Usuario struct {
	nombre     string
	correo     string
	contrasena string
	rol        string // "admin" o "usuario"
}

// NuevoUsuario crea un nuevo usuario validando su correo y contraseña
func NuevoUsuario(nombre, correo, contrasena, rol string) (*Usuario, error) {
	if nombre == "" || correo == "" || contrasena == "" {
		return nil, errors.New("todos los campos son obligatorios")
	}

	if !validarCorreo(correo) {
		return nil, errors.New("formato de correo inválido")
	}

	if !validarContrasena(contrasena) {
		return nil, errors.New("la contraseña debe tener mínimo 6 caracteres, incluyendo letras y números")
	}

	if rol != "admin" && rol != "usuario" {
		return nil, errors.New("rol inválido, debe ser 'admin' o 'usuario'")
	}

	return &Usuario{
		nombre:     nombre,
		correo:     correo,
		contrasena: contrasena,
		rol:        rol,
	}, nil
}

// ======================= MÉTODOS GET =======================

func (u *Usuario) GetNombre() string {
	return u.nombre
}

func (u *Usuario) GetCorreo() string {
	return u.correo
}

func (u *Usuario) GetRol() string {
	return u.rol
}

// ======================= MÉTODOS SET =======================

func (u *Usuario) SetNombre(nombre string) {
	u.nombre = nombre
}

func (u *Usuario) SetCorreo(correo string) error {
	if !validarCorreo(correo) {
		return errors.New("correo inválido")
	}
	u.correo = correo
	return nil
}

func (u *Usuario) SetContrasena(contra string) error {
	if !validarContrasena(contra) {
		return errors.New("contraseña no válida")
	}
	u.contrasena = contra
	return nil
}

func (u *Usuario) SetRol(rol string) error {
	if rol != "admin" && rol != "usuario" {
		return errors.New("rol no permitido")
	}
	u.rol = rol
	return nil
}

// ======================= VALIDACIONES =======================

func validarCorreo(correo string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(correo)
}

func validarContrasena(contra string) bool {
	if len(contra) < 6 {
		return false
	}

	tieneLetra := false
	tieneNumero := false

	for _, c := range contra {
		if c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' {
			tieneLetra = true
		}
		if c >= '0' && c <= '9' {
			tieneNumero = true
		}
	}

	return tieneLetra && tieneNumero
}

