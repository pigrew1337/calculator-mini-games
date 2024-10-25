package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var randomNumber int

func initRandomNumber() {
	rand.Seed(time.Now().UnixNano())
	randomNumber = rand.Intn(1000)
}

func Igra(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/igra.html"))

	if r.Method == http.MethodPost {
		userInputStr := r.FormValue("number")
		userInput, err := strconv.Atoi(userInputStr)
		var message string
		var result string
		var fottersText string

		if err != nil {
			message = "Введите корректное число"
		} else if userInput > randomNumber {
			message = "число меньше"
		} else if userInput < randomNumber {
			message = "число больше"
		} else {
			message = "Поздравляем! Вы угадали число! "
			result = strconv.Itoa(randomNumber)
			fottersText = "И ЭТО ЧИСЛО ПРИДУМАЛ ДМИТРИЙ УТКИН "
			initRandomNumber()
		}

		data := struct {
			Message      string
			Result       string
			FottersTextа string
		}{
			Message:      message,
			Result:       result,
			FottersTextа: fottersText,
		}
		tmpl.Execute(w, data)
		return
	}

	initRandomNumber()
	tmpl.Execute(w, nil)
}
