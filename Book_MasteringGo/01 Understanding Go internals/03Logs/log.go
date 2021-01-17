package main

import (
	"log"
	"os"
)

func main() {
	//Aca le digo que flags quiero usar, separadas por |
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//Write to a file
	// O_APPEND : agregar a los existentes
	// O_CREATE : crear si no existe
	// O_WRONLY : solo escritura
	// 0666 : read & write, para permisos de un archivo en Unix
	file, err := os.OpenFile("Book_MasteringGo/03Logs/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println("This is my log message")

	//Print the exits status only in console
	//log.Fatal("This is my log message 2")

	//Print the runtime panic
	//log.Panic("This is my log message 3")
}
