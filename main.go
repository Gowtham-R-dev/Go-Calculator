package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

)

// Calculation holds data for a single calculation
type Calculation struct {
	FirstNum  float64
	SecondNum float64
	Result    float64
	Operator  string
}

// calculatorHandler handles both GET and POST requests for the calculator
func calculatorHandler(w http.ResponseWriter, r *http.Request) {
	// Define the template path
	tmplPath := filepath.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		// Parse form values
		firstNumStr := r.FormValue("firstNum")
		secondNumStr := r.FormValue("secondNum")
		operator := r.FormValue("operator")

		// Convert string values to float64
		firstNum, err := strconv.ParseFloat(firstNumStr, 64)
		if err != nil {
			http.Error(w, "Invalid first number", http.StatusBadRequest)
			return
		}
		secondNum, err := strconv.ParseFloat(secondNumStr, 64)
		if err != nil {
			http.Error(w, "Invalid second number", http.StatusBadRequest)
			return
		}

		// Perform the calculation
		var result float64
		switch operator {
		case "+":
			result = firstNum + secondNum
		case "-":
			result = firstNum - secondNum
		case "*":
			result = firstNum * secondNum
		case "/":
			if secondNum != 0 {
				result = firstNum / secondNum
			} else {
				http.Error(w, "Cannot divide by zero", http.StatusBadRequest)
				return
			}
		default:
			http.Error(w, "Invalid operator", http.StatusBadRequest)
			return
		}

		// Create a Calculation instance with the results
		calculation := Calculation{
			FirstNum:  firstNum,
			SecondNum: secondNum,
			Operator:  operator,
			Result:    result,
		}

		// Render the template with the calculation result
		if err := tmpl.Execute(w, calculation); err != nil {
			http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Render the template without any calculation if it's a GET request
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", calculatorHandler)
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
