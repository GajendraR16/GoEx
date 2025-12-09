package main

import "fmt"

type Notifier interface {
	Send(message string) error
}

// TODO: Implement these types
type EmailNotifier struct {
	ToAddress string
}

type SMSNotifier struct {
	PhoneNumber string
}

type SlackNotifier struct {
	Channel string
}

// TODO: Implement Send() for each type
func (e EmailNotifier) Send(message string) error {
	fmt.Printf("📧 Email sent to %s: %s\n", e.ToAddress, message)
	return nil
}

func (s SMSNotifier) Send(message string) error {
	fmt.Printf("📱 SMS sent to %s: %s\n", s.PhoneNumber, message)
	return nil
}

func (s SlackNotifier) Send(message string) error {
	fmt.Printf("💬 Slack message sent to %s: %s\n", s.Channel, message)
	return nil
}

// TODO: Write a function that sends to multiple notifiers
func NotifyAll(notifiers []Notifier, message string) {
	for _, v := range notifiers {
		v.Send(message)
	}
}

func main() {
	email := EmailNotifier{ToAddress: "user@example.com"}
	sms := SMSNotifier{PhoneNumber: "+1234567890"}
	slack := SlackNotifier{Channel: "#general"}

	notifiers := []Notifier{email, sms, slack}
	NotifyAll(notifiers, "System alert!")
}
