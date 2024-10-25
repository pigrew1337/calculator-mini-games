package main

import (
	"html/template"
	"math"
	"net/http"
	"strconv"
)

type Result struct {
	Number1   float64
	Number2   float64
	Operation string
	Result    float64
	Error     string
}

func Calculate(w http.ResponseWriter, r *http.Request) {
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
