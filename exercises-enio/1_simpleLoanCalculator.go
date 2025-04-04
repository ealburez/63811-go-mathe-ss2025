package main

import(
	"fmt"
	"math"
)

func monthlyPaymentCalculator(loanAmount float64, annualRate float64, years int) float64{
	var r float64 = annualRate/12.0
	var n float64 = float64(years)*12.0
	return loanAmount * r * (math.Pow(1.0 + r,n)/(math.Pow(1.0 + r,n)-1.0))

}

func main() {
	var monthlyPayment float64
	var loanAmount float64
	var annualRate float64
	var years int

	fmt.Println("Principal Amount:")
	fmt.Scan(&loanAmount)

	fmt.Println("Annual Interest Rate:")
	fmt.Scan(&annualRate)
	
	//convert from percentage to float
	annualRate = annualRate / 100.0

	fmt.Println("Loan term in years:")
	fmt.Scan(&years)

	monthlyPayment = monthlyPaymentCalculator(loanAmount, annualRate, years)

	fmt.Println("Your monthly payment is:", math.Round(monthlyPayment*100)/100)

	THERE IS AN ERROR IN THE MATHS NEED TO CHECK THE EXACT RESULT OF THE INPUT 
	100000
	5
	30
}