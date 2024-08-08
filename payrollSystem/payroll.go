package payrollSystem

// func SimplePayrollSystem() {
// 	reader := bufio.NewReader(os.Stdin)

// 	// Input employee name
// 	fmt.Print("Enter employee name: ")
// 	name, _ := reader.ReadString('\n')
// 	name = strings.TrimSpace(name)

// 	// Input rate per hour
// 	fmt.Print("Enter rate per hour: ")
// 	ratePerHourStr, _ := reader.ReadString('\n')
// 	ratePerHourStr = strings.TrimSpace(ratePerHourStr)
// 	ratePerHour, _ := strconv.ParseFloat(ratePerHourStr, 64)

// 	// Input hours per day
// 	fmt.Print("Enter hours per day: ")
// 	hoursPerDayStr, _ := reader.ReadString('\n')
// 	hoursPerDayStr = strings.TrimSpace(hoursPerDayStr)
// 	hoursPerDay, _ := strconv.ParseFloat(hoursPerDayStr, 64)

// 	// Input number of days worked
// 	fmt.Print("Enter number of days worked: ")
// 	numberOfDaysWorkedStr, _ := reader.ReadString('\n')
// 	numberOfDaysWorkedStr = strings.TrimSpace(numberOfDaysWorkedStr)
// 	numberOfDaysWorked, _ := strconv.Atoi(numberOfDaysWorkedStr)

// 	// Compute gross salary
// 	grossSalary := ratePerHour * hoursPerDay * float64(numberOfDaysWorked)

// 	// Compute deductions
// 	tax := grossSalary * 0.15
// 	philHealth := grossSalary * 0.05
// 	sss := grossSalary * 0.02

// 	totalDeduction := tax + philHealth + sss

// 	// Compute net salary
// 	netSalary := grossSalary - totalDeduction

// 	// Output
// 	fmt.Printf("Employee Name: %s\n", name)
// 	fmt.Printf("Gross Salary: %.2f\n", grossSalary)
// 	fmt.Printf("Tax Deduction: %.2f\n", tax)
// 	fmt.Printf("PhilHealth Deduction: %.2f\n", philHealth)
// 	fmt.Printf("SSS Deduction: %.2f\n", sss)
// 	fmt.Printf("Total Deduction: %.2f\n", totalDeduction)
// 	fmt.Printf("Net Salary: %.2f\n", netSalary)
// }
