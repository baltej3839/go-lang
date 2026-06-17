package main

import "fmt"

type Storage interface {
    Save(data string) error
}

type EmailNotifier struct {
	email string
}

func (e *EmailNotifier) Send(msg string) error {
	fmt.Print("msg sent via email")
	return nil
}

func NewEmailNotifer() *EmailNotifier {
	return &EmailNotifier{}
}

type SMSNotifier struct {
	sms string
}

func NewSMSNotifier() *SMSNotifier {
	return &SMSNotifier{}
}

func (e *SMSNotifier) Send(msg string) error {
	fmt.Print("msg sent via sms")
	return nil
}

type PushNotifier struct {
	pushN string
}

func NewPushNotifier() *PushNotifier {
	return &PushNotifier{}
}

func (e *PushNotifier) Send(msg string) error {
	fmt.Print("msg sent via push service")
	return nil
}

func NewNotifier(notifierMedium string) (Notifier, error) {
	switch notifierMedium {
	case "sms":
		return NewSMSNotifier(), nil 
	case "email":
		return NewEmailNotifer(), nil
	case "push":
		return NewPushNotifier(), nil 
	default:
		return nil, fmt.Errorf("medium not present")
	}

}

func main() {
	notifier, err:=NewNotifier("sms")
	if err!=nil {
		fmt.Print("error is printed for this")		
	} 
	notifier.Send("Hi from Baltej")
}