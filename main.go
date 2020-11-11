package main

import (

	"fmt"
	"time"
)

type proceso struct {
	id    int
	i     chan uint64
	done  chan bool
	
}gi

func (p *proceso) start() {
	var i uint64 = 0
	for {
	
		i++
		p.i <- i
	}

}

func (p* proceso) stop() {
	for {

		select {
		case <-p.i:
			time.Sleep(time.Millisecond * 500)
		case <-p.done: //cierra el canal
			close(p.done)
			return
		}
	}
}

func imprimir(quit chan bool, s []proceso){//mostrar todo
	for {
		select {
		case <-quit:
			return
		default:
			
			for _, j := range s {
				fmt.Print( j.id)
				msg := <-j.i
				fmt.Println(" : ", msg)
				time.Sleep(time.Millisecond * 500)
			}
		}
	}
}


func main() {
	n := 0 //id
	var s []proceso//slice
	var op string
	for {
		
		fmt.Println("1. Agregar proceso \n2. Eliminar Proceso \n3. Mostrar procesos \n0. Salir ")
		fmt.Print("Ingrese opción : ")

		fmt.Scan(&op)
		switch (op) {

			case "0":
				return
			
			case "1":
				
				p := proceso{n, make(chan uint64),make(chan bool)}
				s = append(s, p)
				go p.start()
				go p.stop()
				n++
			case "2":
				var a int
				fmt.Print("ID que quiere eliminar: ")
				fmt.Scan(&a)
				var element *proceso
				for idx, e := range s {
					if e.id == a {
						element = &e
						s = append(s[:idx], s[idx+1:]...)
						break
					}
				}
				element.done <- false
			case "3":
				
				var input string
				quit := make(chan bool)
				go imprimir(quit, s)		
				fmt.Scan(&input)					
				if(input =="2"){
					quit <- true
					break
				}
			}
		}
	
}