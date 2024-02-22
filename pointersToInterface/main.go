package main

import "fmt"

// Notifier is an interface for sending notifications
type Notifier interface {
	Notify(message string)
}

// EmailNotifier is a concrete type implementing Notifier
type EmailNotifier struct {
	EmailAddress string
}

// Notify sends an email notification
func (e EmailNotifier) Notify(message string) {
	fmt.Printf("Sending email to %s: %s\n", e.EmailAddress, message)
}

// SMSNotifier is another concrete type implementing Notifier
type SMSNotifier struct {
	PhoneNumber string
}

// Notify sends an SMS notification
func (s SMSNotifier) Notify(message string) {
	fmt.Printf("Sending SMS to %s: %s\n", s.PhoneNumber, message)
}

// NotifyUser sends a notification to the provided Notifier
func NotifyUser(notifier Notifier, message string) {
	notifier.Notify(message)
}

func main() {
	// Create instances of concrete notifiers
	emailNotifier := EmailNotifier{EmailAddress: "john@example.com"}
	smsNotifier := SMSNotifier{PhoneNumber: "+123456789"}

	// Use NotifyUser function with different notifiers
	NotifyUser(emailNotifier, "Hello via Email")
	NotifyUser(smsNotifier, "Hello via SMS")
}
