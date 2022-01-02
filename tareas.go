package main

import "fmt"

type taskList struct{
	tasks []*task
}

func (t *taskList) agregarAlista( tl *task){
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) eliminarDeList(index int){
	t.tasks = append(t.tasks[:index], t.tasks[index + 1:]... )
}

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
	t1 := &task{
		nombre:      "Completar mi curso de Go",
		descripcion: "Completar mi curso esta semana",
	}

	t2 := &task{
		nombre:      "Completar mi curso de Java",
		descripcion: "Completar mi curso esta semana",
	}

	t3 := &task{
		nombre:      "Completar mi curso de Python",
		descripcion: "Completar mi curso esta semana",
	}
	
	lista := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}

	fmt.Println(lista.tasks[0])
	lista.agregarAlista(t3)
	fmt.Println(len(lista.tasks))
	lista.eliminarDeList(1)
	fmt.Println(len(lista.tasks))
}