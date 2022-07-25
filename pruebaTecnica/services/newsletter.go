package services

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// EmailObject defines email payload data
type EmailObject struct {
	To      string
	Body    string
	Subject string
}

var emailPass = []byte(os.Getenv("MAIL_SECRET"))

// SendMail method to send email to user
func SendMail(subject string, body string, to string, html string, name string) bool {
	fmt.Println(os.Getenv("SENDGRID_API_KEY"))

	from := mail.NewEmail("Just Open it", os.Getenv("SENDGRID_FROM_MAIL"))
	_to := mail.NewEmail(name, to)
	plainTextContent := body
	htmlContent := html
	message := mail.NewSingleEmail(from, subject, _to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	_, err := client.Send(message)
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}
}

func SendMailGo() {

	auth := smtp.PlainAuth("", "raulmar.garca@gmail.com", "1f2g3h4j", "smtp.gmail.com.")
	to := []string{"perajim@gmail.com"} // Array de correos de destino

	// El texto del body con la información del correo puede ser sensible a los espacios.
	body := fmt.Sprintf(`From: "MyAppGo ✉️" <prueba@micorreo.com>
	Subject:%s
	%s`, "Correo desde Go", "Este correo es enviado desde una app go")

	msg := []byte(body)

	error := smtp.SendMail("smtp.proveedor.com:587", auth, "raulmar.garca@gmail.com", to, msg)
	if error != nil {
		fmt.Println("Informamos el error", error)
	}
}
