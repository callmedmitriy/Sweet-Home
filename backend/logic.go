package main

import (
	"log"
	"fmt"
	"math"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

// this is a comment
var settings Settings
var cityList [3]string 


func main() {
	cityList = [3]string{"Ростов","Москва","Питер"}
	settings = Settings{
		City: 10,
		MortgageInterest: 9.5,
		DepositInterest: 6,
		FirstPay: 1000000,
		RoomCount: 2,
		newFlat: true,
		firstBuy: true,
		ownFlat: true,

		RentPayment: 18000,
		Repair: 1000000,
		Inflation: 15, 
	}

	var creditAmount int = 2300000 
	var loanRate float64 = 12
	var n int = 60

	annuityPayment(creditAmount,loanRate,n)
	differentiatedPayment(creditAmount,loanRate,n)

	r := mux.NewRouter()
	r.HandleFunc("/", getHello).Methods(http.MethodGet)
	r.HandleFunc("/city", getCityList).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":3020", r))
}

func getHello(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get hello world"}`))

}

func getCityList(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(settings.City)

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

/*
	Structures
*/

type Settings struct {
	City int `json:"city"`
	MortgageInterest float32 `json:"mortgageInterest"`
	DepositInterest float32 `json:"depositInterest"`
	FirstPay int `json:"firstPay"`
	RoomCount int `json:"roomCount"`
	newFlat bool `json:"newFlat"`
	firstBuy bool `json:"firstBuy"`
	ownFlat bool `json:"ownFlat"`

	RentPayment int `json:"RentPayment"`
	Repair int `json:"repair"`
	Inflation int `json:"inflation"`
	Options *[]Option `json:"options"`
}

type Option struct {
	Time int `json:"time"`
	AnnuitySum int `json:"annuitySum"`
	DifferentiatedySum int `json:"differentiatedySum"`
	DepositSum int `json:"depositSum"`
	MonthlyIncome int `json:"monthlyIncome"`
	Calculation []*monthlyCalculation `json:"calculation"`
}

type monthlyCalculation struct {
	Month int `json:"month"`
	Annuity int `json:"annuity"`
	Differentiated int `json:"differentiated"`
	RentIncome int `json:"rentIncome"`
	RentConsumption int `json:"rentConsumption"`
	Deposit int `json:"deposit"`
	Addition int `json:"addition"`
}