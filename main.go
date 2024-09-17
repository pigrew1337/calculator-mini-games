package main

import (
	"fmt"
	"html/template"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Result struct {
	Number1   float64
	Number2   float64
	Operation string
	Result    float64
	Error     string
}

func main() {
	// статич файл
	fs := http.FileServer(http.Dir("templates/css"))
	// работаем
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.HandleFunc("/", calculate)
	http.HandleFunc("/igra", igra)

	fmt.Println("Сервер запущен, порт 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func calculate(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/calculate.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	num1, err1 := strconv.ParseFloat(r.FormValue("number1"), 64)
	num2, err2 := strconv.ParseFloat(r.FormValue("number2"), 64)
	operation := r.FormValue("operation")
	var result float64
	var errMessage string

	if err1 != nil || err2 != nil {
		errMessage = "Ошибка: введите корректные числа"
	} else {
		switch operation {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "^":
			result = math.Pow(num1, num2)
		case "/":
			if num2 == 0 {
				errMessage = "Ошибка: деление на ноль"
			} else {
				result = num1 / num2
			}
		default:
			errMessage = "Неизвестная операция"
		}
	}

	res := Result{
		Number1:   num1,
		Number2:   num2,
		Operation: operation,
		Result:    result,
		Error:     errMessage,
	}

	tmpl.Execute(w, res)
}

var randomNumber int

func initRandomNumber() {
	rand.Seed(time.Now().UnixNano())
	randomNumber = rand.Intn(1000)
}

func igra(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/igra.html"))

	if r.Method == http.MethodPost {
		userInputStr := r.FormValue("number")
		userInput, err := strconv.Atoi(userInputStr)
		var message string

		if err != nil {
			message = "Введите корректное число"
		} else if userInput > randomNumber {
			message = "число меньше"
		} else if userInput < randomNumber {
			message = "число больше"
		} else {
			message = "Поздравляем! Вы угадали число!"
			initRandomNumber()
		}

		data := struct {
			Message string
		}{
			Message: message,
		}
		tmpl.Execute(w, data)
		return
	}

	initRandomNumber()
	tmpl.Execute(w, nil)
}
