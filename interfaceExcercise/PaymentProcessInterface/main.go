package main

import "fmt"

type PaymentProcesser interface {
	ProcessPayment(int)
}


type order struct {
	orderNumber int
	price       int
}

func CheckOrder(o order) bool {
	fmt.Print("Order is valid")
	return true 
}

type Stripe struct {
	stripeRegNumber int
}

func NewStripeService() *Stripe {
	return &Stripe{stripeRegNumber: 1234}
}

func (s *Stripe) ProcessPayment(money int) {
	fmt.Print("payment process done by stripe")
}


type Razorpay struct {
	RazorpayRegNumber int
}


func NewRazorpayService() *Razorpay {
	return &Razorpay{RazorpayRegNumber: 1234}
}

func (s *Razorpay) ProcessPayment(money int) {
	fmt.Print("payment process done by Razorpay")
}


type PayPal struct {
	PayPalRegNumber int
}

func NewPayPalService() *PayPal {
	return &PayPal{PayPalRegNumber: 1234}
}


func (s *PayPal) ProcessPayment(money int) {
	fmt.Print("payment process done by PayPal")
}


// func OrderService(o order, p PaymentProcesser){
// 	CheckOrder(o)
// 	p.ProcessPayment(o.price)
// }

type OrderService struct {
	paymentProcesser PaymentProcesser
}

func NewOrderService(p PaymentProcesser) *OrderService {
	return &OrderService{paymentProcesser: p}
}

func (s *OrderService) ProcessOrder(o order){
	s.paymentProcesser.ProcessPayment(o.price)
	fmt.Print("Order placed")

}

func main() {
	o:=order{orderNumber: 123, price:999}
	
	var p PaymentProcesser=NewRazorpayService()
	s:=NewOrderService(p)
	s.ProcessOrder(o)

	
	
}