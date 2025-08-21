# Go Project - Sistema de Gestión de Empleados (Práctica)

Este proyecto es un sistema básico de gestión de empleados desarrollado en Go.  
Este proyecto es únicamente para fines de práctica y aprendizaje.

## Funcionalidades implementadas

- Inicio de sesión para usuarios y supervisores.
- Encriptación y desencriptación de IDs de usuario.
- Creación de empleados por parte del supervisor.
- Restauración de contraseña.
- Eliminación automática del archivo `id.txt` al cerrar el sistema (incluyendo cierre por `CTRL+C`).
- Uso de SQLite como base de datos local.

## Estructura principal

- `main.go`: Lógica principal y menú de opciones.
- `db.go`: Inicialización y manejo de la base de datos.
- `encrypt.go`: Funciones para encriptar y desencriptar datos.
- `supervisor.go`: Funciones para supervisores (agregar empleados).
- `id.txt`: Archivo temporal para almacenar el ID encriptado del usuario logueado.

## Notas

- El proyecto está en desarrollo y sirve únicamente como práctica de Go y manejo de bases de datos.
- El código no está optimizado para producción ni incluye validaciones avanzadas de seguridad.
- Puedes ejecutar el proyecto con:
  ```sh
  go run .
  ```

---
**Este proyecto es solo de práctica.**