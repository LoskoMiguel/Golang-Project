package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func supervisorOption() {
	var option int
	fmt.Print("Seleccione una opcion:\n 1. Agregar Empleado \n 2. Ver Empleados \n 3. Salir\n")
	fmt.Scan(&option)

	if option == 1 {
		createEmployee()
	} else if option == 2 {
		fmt.Println("En Proceso...")
	} else if option == 3 {
		fmt.Println("Saliendo del sistema")
		return
	} else {
		fmt.Println("Opcion no valida")
	}
}

func login() {
	InitDB()
	defer DB.Close()

	var username string
	var password string

	fmt.Print("Ingrese Su Usuario: ")
	fmt.Scan(&username)

	fmt.Print("Ingrese Su Contrasena:  ")
	fmt.Scan(&password)

	var id int
	var rol string
	err := DB.QueryRow("SELECT id, rol FROM users WHERE username = ? AND password = ?", username, password).Scan(&id, &rol)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Usuario o Contrasena Incorrecta")
		} else {
			log.Fatal(err)
		}
		return
	}

	key := "1234567890123456"

	enc, err := encrypt(id, key)
	if err != nil {
		panic(err)
	}

	os.WriteFile("id.txt", []byte(enc), 0644)
	fmt.Println("ID Encriptado Guardado En id.txt")
	fmt.Println("Bienvendio Al Sistema")
	if rol == "supervisor" {
		supervisorOption()
	} else {
		fmt.Println("En Proceso...")
	}
}

func restorePassword() {
	InitDB()
	defer DB.Close()

	var username string
	fmt.Print("Ingrese Su Usuario: ")
	fmt.Scan(&username)

	var password string
	err := DB.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Usuario No Encontrado")
		} else {
			log.Fatal(err)
		}
		return
	}

	if password == "no" {
		var newPassword string
		fmt.Print("Ingrese Nueva Contrasena: ")
		fmt.Scan(&newPassword)

		_, err = DB.Exec("UPDATE users SET password = ? WHERE username = ?", newPassword, username)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Contrasena Actualizada Exitosamente")
		}
	} else {
		fmt.Println("La Contrasena Ya Esta Establecida")
	}
}

func deleteIDFile() {
	os.Remove("id.txt")
}

func main() {
	InitDB()
	defer DB.Close()

	defer deleteIDFile()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		deleteIDFile()
		fmt.Println("\nArchivo id.txt borrado. Saliendo del sistema.")
		os.Exit(0)
	}()

	fmt.Println("Bienvenido al sistema de gestion de empleados")

	var option int
	fmt.Print("Seleccione una opcion:\n 1. Iniciar Sesion \n 2. Restaurar Contrasena \n 3. Salir\n")
	fmt.Scan(&option)

	if option == 1 {
		login()
	} else if option == 2 {
		restorePassword()
	} else if option == 3 {
		fmt.Println("Saliendo del sistema")
		return
	} else {
		fmt.Println("Opcion no valida")
	}
}
