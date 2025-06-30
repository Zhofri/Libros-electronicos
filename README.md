Sistema de Gestión de Libros Electrónicos

📚 Información del Proyecto

Estudiante: Zhofri Joel Guaman Quichimbo

Universidad: Universidad Internacional Del Ecuador (UIDE)

Materia: Programación Orientada a Objetos

Fecha: 29 de Junio 2025

Lenguaje: Go (Golang)

Framework: Gin

🎯 Objetivo del Programa

Desarrollar un sistema web completo de gestión de libros electrónicos que demuestre la implementación de servicios web, manejo de bases de datos, autenticación JWT, y arquitectura MVC utilizando Go como lenguaje principal. El sistema permite a los usuarios registrarse, autenticarse, comprar libros, solicitar préstamos, y a los administradores gestionar el catálogo completo.

🚀 Funcionalidades Principales

Para Usuarios Regulares:

✅ Registro y autenticación segura

✅ Visualización del catálogo de libros

✅ Compra de libros electrónicos

✅ Sistema de préstamos temporales

✅ Biblioteca personal de libros adquiridos

✅ Gestión de perfil de usuario

Para Administradores:

✅ Panel de administración completo

✅ CRUD completo de libros (Crear, Leer, Actualizar, Eliminar)

✅ Gestión de portadas de libros

✅ Control de disponibilidad y precios

✅ Supervisión de transacciones

🛠️ Stack Tecnológico

Backend: Go (Golang) con framework Gin

Base de Datos: MySQL con ORM GORM

Autenticación: JWT (JSON Web Tokens)

Frontend: HTML Templates con Gin

Encriptación: bcrypt para passwords

Configuración: Variables de entorno (.env)


📁 Estructura del Proyecto
LIBRO_ELECTRONICO/

├── controllers/           # Controladores de rutas y lógica de negocio

│   ├── actualizar_libro.go

│   ├── agregar_libro.go

│   ├── cerrar_sesion.go

│   ├── comprar_libro.go

│   ├── datos_libro_editar.go

│   ├── eliminar_libro.go

│   ├── iniciar_sesion.go

│   ├── libros_admin.go

│   ├── libros_adquiridos.go

│   ├── libros.go

│   ├── prestamo_libro.go

│   └── registro.go


├── db/                    # Configuración de base de datos

│   └── db.go


├── middlewares/           # Funciones de validación y seguridad

│   ├── isAdmin.go
│   └── validar_usuario.go


├── models/                # Modelos de datos

│   ├── libro.go

│   ├── transaccion.go

│   └── usuario.go


├── static/                # Archivos estáticos

│   └── portadas/          # Portadas de libros

├── utils/                 # Utilidades


│   ├── generarJWT.go

│   └── password.go


├── views/                 # Templates HTML

│   ├── crear_libro.html

│   ├── editar_libro.html

│   ├── index.html

│   ├── iniciar_sesion.html

│   ├── inicio.html

│   ├── libros_admin.html

│   ├── libros_adquiridos.html

│   ├── registro.html

│   └── transacciones.html


├── main.go               # Punto de entrada principal


├── .env                  # Variables de entorno


└── go.mod               # Dependencias del proyecto


🔧 Los 8 Servicios Web Implementados

Servicio de Registro - POST /registro

Servicio de Autenticación - POST /login

Servicio de Catálogo - GET /autenticado/inicio

Servicio de Compras - POST /autenticado/inicio/libros/comprar/:id

Servicio de Préstamos - POST /autenticado/inicio/libros/prestamo/:id

Servicio de Biblioteca Personal - GET /autenticado/inicio/libros/adquiridos

Servicio de Administración - GET /autenticado/admin/libros

Servicio CRUD de Libros - POST/GET/PUT/DELETE /autenticado/admin/libros/*

🔐 Seguridad Implementada

Autenticación JWT: Tokens seguros con expiración de 24 horas

Encriptación de Passwords: bcrypt con salt automático

Middlewares de Validación: Control de acceso por niveles

Validación de Archivos: Subida segura de portadas

Variables de Entorno: Configuración sensible protegida

📊 Modelos de Datos

Usuario

gotype Usuario struct {

    gorm.Model

    Nombre        string

    Email         string        `gorm:"unique;not null"`

    Password      string        `json:"-"`

    IsAdmin       bool          `gorm:"default:false"`

    Transacciones []Transaccion

}

Libro

gotype Libro struct {

    gorm.Model

    Titulo      string  `gorm:"not null"`

    Autor       string  `gorm:"not null"`

    Descripcion string

    Categoria   string  `gorm:"not null"`

    Disponible  bool    `gorm:"default:true"`

    Precio      float64

    Portada     string

}

Transacción

gotype Transaccion struct {

    gorm.Model

    Tipo        string     // "compra" o "prestamo"

    LibroID     uint

    UsuarioID   uint

    FechaInicio time.Time

    FechaFin    *time.Time

    Libro       Libro

    Usuario     Usuario

}

🚀 Instalación y Ejecución

envDB_USER=admin

DB_PASSWORD=admin

DB_HOST=localhost

DB_PORT=3306

DB_NAME=libros_electronicos

JWT_SECRET=tu_secret_key

Ejecutar la aplicación

bashgo run main.go

Acceder a la aplicación

http://localhost:8080

🧪 Testing con Postman


🔄 Integración de Conocimientos

El proyecto integra los siguientes temas de la materia:

Manejo de Funciones y Paquetes: Organización modular del código
Arrays, Slices, Maps: Manejo de colecciones de datos
Structs: Definición de modelos de datos
Métodos y Funciones: Implementación de lógica de negocio
Constructores: Inicialización de estructuras
Encapsulación: Control de acceso a datos
Interfaces: Abstracci

📈 Serialización JSON

Todos los modelos implementan serialización JSON automática:

gotype Libro struct {

    Titulo string `json:"titulo"`

    Autor  string `json:"autor"`

    // ...

}

🌟 Características Destacadas


Arquitectura MVC: Separación clara de responsabilidades

Middleware en Capas: Seguridad escalonada

Manejo de Archivos: Subida segura de portadas

Dual Mode: Compra permanente y préstamo temporal

Panel Admin: Gestión completa del sistema

Responsive Design: Compatible con diferentes dispositivos

🎯 Casos de Uso

Usuario Regular: Registro → Login → Explorar catálogo → Comprar/Prestar libros

Administrador: Login → Gestionar libros → Subir portadas → Controlar inventario

📝 Notas de Desarrollo

El sistema utiliza automigración de GORM para crear las tablas

Las contraseñas se encriptan automáticamente con bcrypt

Los JWT tokens incluyen información de rol para autorización

El sistema maneja tanto compras permanentes como préstamos temporales

🤝 Contribución

Este es un proyecto educativo individual desarrollado como parte del curso de Programacion Orientada A Objetos.

📄 Licencia

Este proyecto es para fines educativos únicamente.
