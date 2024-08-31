package responses

import (
	"io"
	"net/smtp"
	"os"
	"strings"
)

func SendToMail(to, subject, body string) {
	email_config_file, _ := os.Open("email")
	email_config_str, _ := io.ReadAll(email_config_file)
	email_config := strings.Split(string(email_config_str), ",")
	host := email_config[0]
	user := email_config[1]
	password := email_config[2]
	email_config_file.Close()

	auth := smtp.PlainAuth("", user, password, host)
	content_type := "Content-Type: text/html; charset=UTF-8"
	msg := []byte("To: " + to + "\r\nFrom: MetaICP Admin<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	_ = smtp.SendMail(host+":25", auth, user, send_to, msg)
}

func GetMail() string {
	email_config_file, _ := os.Open("email")
	email_config_str, _ := io.ReadAll(email_config_file)
	email_config := strings.Split(string(email_config_str), ",")
	user := email_config[1]
	email_config_file.Close()

	return user
}
