package main

import (
	"fmt"
	"os"

	"github.com/polds/imgbase64"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
	from := mail.NewEmail("Picletter Bot", os.Getenv("BOT_EMAIL"))
	subject := "Picletter, Your daily picture newsletter!"
	to := mail.NewEmail("Rahul", os.Getenv("RECEIVER_EMAIL"))
	// download image from the Unsplash API
	img := imgbase64.FromRemote("https://source.unsplash.com/daily/")
	content := mail.NewContent("text/html", fmt.Sprintf("<img src=\"%v\">", img))
	// compose email
	m := mail.NewV3MailInit(from, subject, to, content)
	request := sendgrid.GetRequest(os.Getenv("SG_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	// send email
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
