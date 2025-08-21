package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func createEmployee() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Ingrese el nombre del empleado: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	username := strings.ReplaceAll(name, " ", "")

	data, _ := os.ReadFile("id.txt")
	decStr, err := decrypt(string(data), "1234567890123456")
	if err != nil {
		panic(err)
	}

	idBoss, err := strconv.Atoi(decStr)
	if err != nil {
		panic(err)
	}

	_, err = DB.Exec("INSERT INTO users (name, username, password, rol, id_boss) VALUES (?, ?, ?, ?, ?)", name, username, "no", "empleado", idBoss)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Empleado agregado exitosamente")
	}
}
