package main

import "fmt"

type Message struct {
	Recipient string
	Text      string
}

func getMessageText(m Message) string {
	return fmt.Sprintf(`To: %v Message: %v`, &m.Recipient, &m.Text)
}

func main() {
	m := Message{
		Recipient: "hassan",
		Text:      "hello g",
	}
	getMessageText(m)
}
