package main

import "fmt"

type PaymentStrategy interface {
	Pay(amount float64) string
}

type CreditCardPayment struct{}

func (c *CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paying %.2f with credit card", amount)
}

type PayPalPayment struct{}

func (p *PayPalPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paying %.2f with paypal", amount)
}

type PaymentContext struct {
	strategy PaymentStrategy
}

func (pc *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	pc.strategy = strategy
}

func (pc *PaymentContext) MakePayment(amount float64) string {
	return pc.strategy.Pay(amount)
}

func main() {
	paymentContext := PaymentContext{}

	creditCard := CreditCardPayment{}
	paypal := PayPalPayment{}

	paymentContext.SetStrategy(&creditCard)
	fmt.Println(paymentContext.MakePayment(100.00))

	paymentContext.SetStrategy(&paypal)
	fmt.Println(paymentContext.MakePayment(50.00))
}
