package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getHealthStatus(bmi float64) string {
    var healthStatus string

    if bmi < 18.5 {
        healthStatus = "Underweight"
    } else if bmi >= 18.5 && bmi <= 24.9 {
        healthStatus = "Normal Weight"
    } else if bmi >= 25.0 && bmi <= 29.9 {
        healthStatus = "Overweight"
    } else if bmi >= 30.0 && bmi <= 34.9 {
        healthStatus = "Obese Class 1"
    } else if bmi >= 35.0 && bmi <= 39.9 {
        healthStatus = "Obese Class 2"
    } else {
        healthStatus = "Obese Class 3"
    }

    return healthStatus
}

func main() {
	const bmiFormulaMsg string = "Formula used: BMI = weight (kg) / (height (m) * height (m))"

	var name, healthStatus string
	var weight, height, bmi float64

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	nameStr, _ := reader.ReadString('\n')
	name = strings.TrimSpace(nameStr)

	fmt.Print("Enter your weight in kg: ")
	weightStr, _ := reader.ReadString('\n')
	weight, _ = strconv.ParseFloat(strings.TrimSpace(weightStr), 64)

	fmt.Print("Enter your height in meters: ")
	heightStr, _ := reader.ReadString('\n')
	height, _ = strconv.ParseFloat(strings.TrimSpace(heightStr), 64)

	bmi = weight / (height * height)
  healthStatus = getHealthStatus(bmi)

	fmt.Println()
	fmt.Printf("Welcome, %s! Let's calculate your BMI.\n", name)
	fmt.Println(bmiFormulaMsg)

	fmt.Println()
	fmt.Printf("Your BMI is: %.2f\n", bmi)
  fmt.Printf("Health Status: %s\n", healthStatus)
}
