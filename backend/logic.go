package main

import (
	"fmt"
	"math"
)

// this is a comment

func main() {
	var creditAmount int = 2300000 
	var loanRate float64 = 12
	var n int = 60

	annuityPayment(creditAmount,loanRate,n)
	differentiatedPayment(creditAmount,loanRate,n)
	//fmt.Println(test(2,3))
}

func annuityPayment(creditAmount int,loanRate float64, n int) float64 {
	// creditAmount - сумма кредита
	// loanRate	- ставка по кредиту (годовая)
	// n - срок кредита (месяцы)

	// l - месячная процентная ставка
	var l float64 = loanRate / 12 / 100

	// ежемесячный платеж
	var monthlyPayment float64 = Round(float64(creditAmount) * ( l * math.Pow(( 1 + l ), float64(n))) / (math.Pow(( 1 + l ), float64(n)) - 1), 2)
	var total float64 = Round(monthlyPayment * float64(n), 2)
	
	fmt.Println("Аннуитетный рассчет:")
	fmt.Println("Ежемесячный платеж:",monthlyPayment)
	fmt.Println("Выплата прцентов:",Round(total - float64(creditAmount), 2))
	fmt.Println("Общая сумма:",int(total))

	// общая сумма выплаты
	return total
}

func differentiatedPayment(creditAmount int,loanRate float64, n int) float64 {
	// creditAmount - сумма кредита
	// loanRate	- ставка по кредиту (годовая)
	// n - срок кредита (месяцы)

	var total, monthlyPayment float64
	var tmpCreditAmount float64 = float64(creditAmount)

	// l - месячная процентная ставка
	var l float64 = loanRate / 12 / 100
	// основной платеж и проценты соответственно
	var b,p float64 

	for i := 1; i <= n; i++ {
		b = Round((float64(creditAmount) / float64(n)), 2)
		p = Round(tmpCreditAmount * l, 2)

		monthlyPayment = b + p
		tmpCreditAmount -= b
		total += monthlyPayment  
	}

	fmt.Println("Диффиренцированный рассчет:")
	fmt.Println("Выплата прцентов:",Round(total - float64(creditAmount),2))
	fmt.Println("Общая сумма:",int(total))

	return total

}

func test(a float64, b float64) float64{
	return Round(a/b,3)
}

func Round(x float64, prec int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}