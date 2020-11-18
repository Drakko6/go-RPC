package main

import (
	// "errors"
	"fmt"
	"net"
	"net/rpc"
)


var Materias map[string]map[string]float64
var Alumnos map[string]map[string]float64


type Server struct{
	
}
	
	
type DatosCrear struct{

	Alumno string
	Materia string
	Cal float64
}


func (this *Server) AgregarCalMateria(crear DatosCrear, respuesta *string) error {
	
	
	// creación de un alumno - calificacion
	alumnoNuevo := make(map[string]float64)
	alumnoNuevo[crear.Alumno] = crear.Cal
  
	var existe bool
	//comprobar si existe la materia
	for mat := range Materias {
		if mat == crear.Materia {
			Materias[crear.Materia][crear.Alumno] = crear.Cal
			
			existe = true
			break
		}
	}
	if !existe{
		
			Materias[crear.Materia] = alumnoNuevo
	}


	// creacion materia - calificacion  
	materiaNueva := make(map[string]float64)
	materiaNueva[crear.Materia] = crear.Cal
  
	// creación de un alumno
	var existe2 bool
	//comprobar si existe la materia
	for alum := range Alumnos {
		if alum == crear.Alumno {
			Alumnos[crear.Alumno][crear.Materia] = crear.Cal
			existe2 = true
			break
		}
	}
	if !existe2{
		
			Alumnos[crear.Alumno] = materiaNueva
	}

	return nil
}

func (this *Server) MostrarMapMaterias(i int64, reply *string) error {

	fmt.Println(Materias)
	fmt.Println(Alumnos)

	*reply = "hOLA"
	return nil
}


func (this *Server) PromedioAlumno(alumno string, reply *float64) error {

	
	var suma float64
	var cont float64
	for _, calificacion := range Alumnos[alumno] {
		
		suma += calificacion
		cont++
	}


	*reply = suma/cont
	return nil
}


func (this *Server) PromedioGeneral(vacio string, reply *float64) error {

	
	var suma float64
	var cont float64
	for alumno := range Alumnos {

		for _, calificacion := range Alumnos[alumno] {
		
			suma += calificacion
			cont++
		}
		
	}


	*reply = suma/cont
	return nil
}


func (this *Server) PromedioMateria(materia string, reply *float64) error {

	
	var suma float64
	var cont float64
	for _, calificacion := range Materias[materia] {
		
		suma += calificacion
		cont++
	}

	*reply = suma/cont
	return nil
}




func server() {
	
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {

	Materias = make(map[string]map[string]float64)
	Alumnos = make(map[string]map[string]float64)

	go server()

	var input string
	fmt.Scanln(&input)
}
