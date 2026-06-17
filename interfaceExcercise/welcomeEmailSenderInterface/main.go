package main

import "fmt"

type WelcomeEmailSender interface{
	SendWelcomeEmail(u UserDetails) error
}

type UserDetails struct {
	email    string
	password string
}

type ProviderA struct {
	providerA string
}

type ProviderB struct {
	providerB string
}

type ProviderC struct {
	providerC string
}

func NewProviderA() ProviderA {
	return ProviderA{providerA: "providerA"}
}

func NewProviderB() *ProviderB {
	return &ProviderB{providerB: "providerB"}
}

func NewProviderC() *ProviderC {
	return &ProviderC{providerC: "providerC"}
}

func (p ProviderA) SendWelcomeEmail(u UserDetails) error {
	fmt.Print(u.email, "welcome email sent from provider A")
	return nil
}

func (p *ProviderB) SendWelcomeEmail(u UserDetails) error {
	fmt.Print(u.email, "welcome email sent from provider B")
	return nil
}

func (p *ProviderC) SendWelcomeEmail(u UserDetails) error {
	fmt.Print(u.email, "welcome email sent from provider C")
	return nil
}


func RegiserUser(u UserDetails, w WelcomeEmailSender) {
	fmt.Print("user is registered", u.email)
	w.SendWelcomeEmail(u)
}





// func SignUpService() SignUper {

// }



func main() {
	var userDetails =UserDetails{email: "asd@gmail.com", password: "12323"}
	var w WelcomeEmailSender=NewProviderA()
	RegiserUser(userDetails, w)
}

// User registers
// when sign up, send welcome email

// What responsibility should belong to the email provider?
// to register a user and send welcome email, to have these methods being with there types

// What responsibility should belong to the registration service?
// register a user and send welcome email

// What interface would you create?
// register and sendWelcome email as methods with SignUper as name of interface

// Why is an interface useful here?
// same methods need to be applied by multiple providers

// What problem would occur if the registration service directly depended on Provider A?
// would have to write code for different providers concretely