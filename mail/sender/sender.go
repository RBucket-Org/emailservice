package sender

import (
	"bytes"
	"html/template"
	"path/filepath"

	"github.com/mailgun/mailgun-go"
)

//TemplateData : provide the structure of data to the given email
type TemplateData struct {
	Name  string
	URL   string
	Site  string
	From  string
	Email string
}

//ParseTemplate : this is the method that parse the html file
func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)

	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	msgBody := buf.String()

	return msgBody, nil
}

//SendEmailVerification : this method sends the email attach with template to verify the email
func SendEmailVerification(domain string, apiKey string, template *TemplateData, templateFileName string, rootdir string) (string, error) {
	mg := mailgun.NewMailgun(domain, apiKey)

	//parse the template
	msgBody, err := ParseTemplate(filepath.Join(rootdir, templateFileName), template)
	if err != nil {
		return "nil", err
	}

	//generate the new message by taking the from, subject, email, body will be empty
	m := mg.NewMessage(
		template.From,
		"Verify Your Email",
		"",
		template.Email,
	)

	//set the HTML
	m.SetHtml(msgBody)

	//Send the message
	_, id, err := mg.Send(m)
	return id, err
}

//SendResetPassword : send the reset password mail to the semder
func SendResetPassword(domain string, apiKey string, template TemplateData, templateFileName string, rootdir string) (string, error) {
	mg := mailgun.NewMailgun(domain, apiKey)

	//parse the template
	msgBody, err := ParseTemplate(filepath.Join(rootdir, templateFileName), template)
	if err != nil {
		return "nil", err
	}

	//generate the new message by taking the from, subject, email, body will be empty
	m := mg.NewMessage(
		template.From,
		"Reset Your Password",
		"",
		template.Email,
	)

	//set the HTML
	m.SetHtml(msgBody)

	//Send the message
	_, id, err := mg.Send(m)
	return id, err
}
