package main

import (
	"fmt"
	"net/rpc"
)


type DatosCrear struct{

	Alumno string
	Materia string
	Cal float64
}

func client() {
	c, err := rpc.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var op int64
	for {


		fmt.Println("1) Agregar calificación de una materia")
		fmt.Println("2) Mostrar promedio de un alumno")
		fmt.Println("3) Mostrar promedio general")
		fmt.Println("4) Mostrar el promedio de una materia")
		fmt.Println("0) Exit")
		fmt.Scanln(&op)

		switch op {
		case 1:

			var alumno string
			fmt.Print("Alumno: ")
			fmt.Scanln(&alumno)

			var materia string
			fmt.Print("Materia: ")
			fmt.Scanln(&materia)

			var cal float64
			fmt.Print("Calificación: ")
			fmt.Scanln(&cal)

			crear := DatosCrear{alumno, materia, cal}
		

			var result string
			err = c.Call("Server.AgregarCalMateria", crear, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result)
			}
		case 2:
			var alumno string
			fmt.Print("Alumno: ")
			fmt.Scanln(&alumno)

			var result float64
			err = c.Call("Server.PromedioAlumno", alumno, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("El promedio de", alumno, "es: ", result)
			}
			
		case 3:

			var vacio string
			var result float64
			err = c.Call("Server.PromedioGeneral", vacio, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("El promedio general es: ", result)
			}
			

		case 4:

			var materia string
			fmt.Print("Materia: ")
			fmt.Scanln(&materia)

			var result float64
			err = c.Call("Server.PromedioMateria", materia, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("El promedio de", materia, "es: ", result)
			}

			
		case 0:
			return
		}
	}
}

func main() {
	client()
}
