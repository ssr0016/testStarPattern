package payrollSystem

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func SimplePayrollWithGui() {
	// Create a new Fyne app
	myApp := app.New()
	myWindow := myApp.NewWindow("Payroll System")

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
	nameLabel := widget.NewLabel("Employee Name: ")
	grossSalaryLabel := widget.NewLabel("Gross Salary: ")
	taxLabel := widget.NewLabel("Tax Deduction: ")
	philHealthLabel := widget.NewLabel("PhilHealth Deduction: ")
	sssLabel := widget.NewLabel("SSS Deduction: ")
	totalDeductionLabel := widget.NewLabel("Total Deduction: ")
	netSalaryLabel := widget.NewLabel("Net Salary: ")

	// Create a button and its event handler
	calculateButton := widget.NewButton("Calculate", func() {
		// Get input values
		name := nameEntry.Text
		ratePerHour, _ := strconv.ParseFloat(ratePerHourEntry.Text, 64)
		hoursPerDay, _ := strconv.ParseFloat(hoursPerDayEntry.Text, 64)
		numberOfDaysWorked, _ := strconv.Atoi(numberOfDaysWorkedEntry.Text)

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
		nameLabel.SetText(fmt.Sprintf("Employee Name: %s", name))
		grossSalaryLabel.SetText(fmt.Sprintf("Gross Salary: %.2f", grossSalary))
		taxLabel.SetText(fmt.Sprintf("Tax Deduction: %.2f", tax))
		philHealthLabel.SetText(fmt.Sprintf("PhilHealth Deduction: %.2f", philHealth))
		sssLabel.SetText(fmt.Sprintf("SSS Deduction: %.2f", sss))
		totalDeductionLabel.SetText(fmt.Sprintf("Total Deduction: %.2f", totalDeduction))
		netSalaryLabel.SetText(fmt.Sprintf("Net Salary: %.2f", netSalary))
	})

	// Layout the input fields, button, and output labels
	myWindow.SetContent(container.NewVBox(
		nameEntry,
		ratePerHourEntry,
		hoursPerDayEntry,
		numberOfDaysWorkedEntry,
		calculateButton,
		nameLabel,
		grossSalaryLabel,
		taxLabel,
		philHealthLabel,
		sssLabel,
		totalDeductionLabel,
		netSalaryLabel,
	))

	// Show and run the app
	myWindow.ShowAndRun()
}
