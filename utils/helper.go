package utils

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"text/template"

	"gopkg.in/gomail.v2"
)

// WriteHeader ...
func (w *ResponseWriterWithLog) WriteHeader(status int) {
	w.Status = status
	w.ResponseWriter.WriteHeader(status)
}

// Write ...
func (w *ResponseWriterWithLog) Write(b []byte) (int, error) {
	w.Body = b
	if w.Status == 0 {
		w.Status = 200
	}
	n, err := w.ResponseWriter.Write(b)
	w.Length += n
	return n, err
}

// UserEmail : ..
type UserEmail struct {
	Name    string
	Email   string
	AuthURL string
}

// SendEmail ..
func SendEmail(token string, email string) bool {
	m := gomail.NewMessage()
	// TODO: read sender email from app config
	m.SetHeader("From", "example@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i><p>Token is: http://localhost:5000/api/v1/users/tokenUser/"+token+"</p>!")
	// m.Attach("/home/Alex/lolcat.jpg")
	fmt.Println("sending email.")
	d := gomail.NewDialer("smtp.gmail.com", 587, "example@gmail.com", "password")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("sending email failed.")
		fmt.Println(err.Error())
		return false
	}
	return true

}

// SendAuthEmail : ..
func (i UserEmail) SendAuthEmail() bool {
	var (
		_, b, _, _ = runtime.Caller(0)
		basepath   = filepath.Dir(b)
	)
	fmt.Println(basepath)
	t := template.New("email.html")
	pt, err := t.ParseFiles("email.html")
	if err != nil {
		log.Println("cannot parse html file ", err.Error())
		return false
	}
	var tpl bytes.Buffer
	if err := pt.Execute(&tpl, i); err != nil {
		log.Println("cannot write data to html file. " + err.Error())
		return false
	}
	tplData := tpl.String()
	m := gomail.NewMessage()
	m.SetHeader("From", "Imran Sarder<example@gmail.com>")
	m.SetHeader("To", i.Email)
	m.SetHeader("")
	m.SetHeader("Subject", "Account Activation")
	m.SetBody("text/html", tplData)
	sent := gomail.NewDialer("smtp.gmail.com", 587, "example@gmail.com", "password")
	if err := sent.DialAndSend(m); err != nil {
		log.Println("cannot send email. " + err.Error())
		return false
	}
	log.Println("email send successfully.")
	return true
}
