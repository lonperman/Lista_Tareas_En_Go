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

func (t *taskList) imprimirLista(){
	for _, tarea := range t.tasks {
		fmt.Println("Nombre:", tarea.nombre)
		fmt.Println("Descripcion:", tarea.descripcion)
	}
}

func (t *taskList) imprimirListaCompletados(){
	for _, tarea := range t.tasks {
		if tarea.completado {
			fmt.Println("Nombre:", tarea.nombre)
			fmt.Println("Descripcion:", tarea.descripcion)
		}
	}
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
	lista.agregarAlista(t3)
	lista.imprimirLista()	
	lista.tasks[0].marcarCompletada()
	fmt.Println("Tareas Completadas")
	lista.imprimirListaCompletados()
	
}