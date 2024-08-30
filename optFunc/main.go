package main

import "fmt"

type Notification struct {
	UserID    string
	Message   string
	SendEmail bool
	SendSMS   bool
	Priority  string
}

type OptFunc func(*Notification)

func WithSMS() OptFunc {
	return func(n *Notification) {
		n.SendSMS = true
	}
}

func WithEmail() OptFunc {
	return func(n *Notification) {
		n.SendEmail = true
	}
}

func WithPriority(priority string) OptFunc {
	return func(n *Notification) {
		n.Priority = priority
	}
}

func SendNotification(userID string, msg string, opts ...OptFunc) {
	// Создаем уведомление с базовыми параметрами
	notification := &Notification{
		UserID:  userID,
		Message: msg,
	}
	// Применяем опции
	for _, opt := range opts {
		opt(notification)
	}

	printNotification(notification)
}

func printNotification(notification *Notification) {
	fmt.Printf("Sending notification to user: %s\n", notification.UserID)
	fmt.Printf("Message: %s\n", notification.Message)

	if notification.SendEmail {
		fmt.Println("Sending Email...")
	}
	if notification.SendSMS {
		fmt.Println("Sending SMS...")
	}
	if notification.Priority != "" {
		fmt.Printf("Priority: %s\n", notification.Priority)
	}

	fmt.Println("Notification sent successfully!")
	fmt.Println("--------")
}

func main() {
	// Вызов без дополнительных опций
	SendNotification("user123", "Hello, world!")

	// Вызов с отправкой Email
	SendNotification("user456", "Important update!", WithEmail())

	// Вызов с отправкой SMS и высоким приоритетом
	SendNotification("user789", "Urgent update!", WithSMS(), WithPriority("high"))
}
