package util

import (
	"fmt"
	"net/smtp"
)

func SendMailMock() {
	fmt.Println("Эмуляция отправки сообщения о смене IP-адреса")
}

func SendMail(receiver string) {
	// Первый попавшийся пример с интернета
	from := "from@gmail.com"
	password := "<Email Password>"
	to := []string{receiver}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	message := []byte("Ваш IP-адрес сменился!")
	
	auth := smtp.PlainAuth("", from, password, smtpHost)
	
	err := smtp.SendMail(smtpHost + ":" + smtpPort, auth, from, to, message)
	if err != nil {
	  fmt.Println(err)
	  return
	}
	fmt.Println("Сообщение о смене IP-адреса отправлено!")
  }