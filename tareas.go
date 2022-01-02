package main

import "fmt"

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t *task) marcarCompletada(){
	t.completado = true
}

func (t *task) actualizarDescripcion(descripcion string){
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string){
	t.nombre = nombre
}

func main() {
	t := &task{
		nombre:      "Completar mi curso de Go",
		descripcion: "Completar mi curso esta semana",
	}
	fmt.Printf("%+v\n", t)
	t.marcarCompletada()
	t.actualizarNombre("Finalizar mi curso")
	t.actualizarDescripcion("Completar mi curso de manera rapida")
	fmt.Printf("%+v\n", t)
}