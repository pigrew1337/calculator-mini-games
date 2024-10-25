package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("templates/css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.HandleFunc("/", Regist)
	http.HandleFunc("/calculate", Calculate)
	http.HandleFunc("/igra", Igra)
	fmt.Println("Сервер запущен, порт 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
