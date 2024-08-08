package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create a new Fyne app
	myApp := app.New()
	myWindow := myApp.NewWindow("Payroll System")

	// Create labels for input fields
	nameLabel := widget.NewLabel("Employee Name:")
	ratePerHourLabel := widget.NewLabel("Rate per Hour:")
	hoursPerDayLabel := widget.NewLabel("Hours per Day:")
	numberOfDaysWorkedLabel := widget.NewLabel("Number of Days Worked:")

	// Create input fields
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Enter employee name")

	ratePerHourEntry := widget.NewEntry()
	ratePerHourEntry.SetPlaceHolder("Enter rate per hour")

	hoursPerDayEntry := widget.NewEntry()
	hoursPerDayEntry.SetPlaceHolder("Enter hours per day")

	numberOfDaysWorkedEntry := widget.NewEntry()
	numberOfDaysWorkedEntry.SetPlaceHolder("Enter number of days worked")

	// Create output labels
	resultNameLabel := widget.NewLabel("Employee Name: ")
	resultGrossSalaryLabel := widget.NewLabel("Gross Salary: ")
	resultTaxLabel := widget.NewLabel("Tax Deduction: ")
	resultPhilHealthLabel := widget.NewLabel("PhilHealth Deduction: ")
	resultSssLabel := widget.NewLabel("SSS Deduction: ")
	resultTotalDeductionLabel := widget.NewLabel("Total Deduction: ")
	resultNetSalaryLabel := widget.NewLabel("Net Salary: ")

	// Create a button and its event handler
	calculateButton := widget.NewButton("Calculate", func() {
		// Get input values
		name := nameEntry.Text
		ratePerHour, err1 := strconv.ParseFloat(ratePerHourEntry.Text, 64)
		hoursPerDay, err2 := strconv.ParseFloat(hoursPerDayEntry.Text, 64)
		numberOfDaysWorked, err3 := strconv.Atoi(numberOfDaysWorkedEntry.Text)

		if err1 != nil || err2 != nil || err3 != nil {
			resultNameLabel.SetText("Please enter valid numbers for rate per hour, hours per day, and number of days worked")
			return
		}

		// Compute gross salary
		grossSalary := ratePerHour * hoursPerDay * float64(numberOfDaysWorked)

		// Compute deductions
		tax := grossSalary * 0.15
		philHealth := grossSalary * 0.05
		sss := grossSalary * 0.02

		totalDeduction := tax + philHealth + sss

		// Compute net salary
		netSalary := grossSalary - totalDeduction

		// Update output labels
		resultNameLabel.SetText(fmt.Sprintf("Employee Name: %s", name))
		resultGrossSalaryLabel.SetText(fmt.Sprintf("Gross Salary: %.2f", grossSalary))
		resultTaxLabel.SetText(fmt.Sprintf("Tax Deduction: %.2f", tax))
		resultPhilHealthLabel.SetText(fmt.Sprintf("PhilHealth Deduction: %.2f", philHealth))
		resultSssLabel.SetText(fmt.Sprintf("SSS Deduction: %.2f", sss))
		resultTotalDeductionLabel.SetText(fmt.Sprintf("Total Deduction: %.2f", totalDeduction))
		resultNetSalaryLabel.SetText(fmt.Sprintf("Net Salary: %.2f", netSalary))
	})

	// Layout the input fields with labels, button, and output labels
	myWindow.SetContent(container.NewVBox(
		nameLabel, nameEntry,
		ratePerHourLabel, ratePerHourEntry,
		hoursPerDayLabel, hoursPerDayEntry,
		numberOfDaysWorkedLabel, numberOfDaysWorkedEntry,
		calculateButton,
		resultNameLabel,
		resultGrossSalaryLabel,
		resultTaxLabel,
		resultPhilHealthLabel,
		resultSssLabel,
		resultTotalDeductionLabel,
		resultNetSalaryLabel,
	))

	// Display and run the application
	myWindow.Resize(fyne.NewSize(400, 400)) // Set a default size
	myWindow.ShowAndRun()
}
