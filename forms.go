// index.go
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/smtp"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func sendHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	message := r.FormValue("message")
	log.Printf("Email body: %s", name)
	log.Printf("Email body: %s", email)
	log.Printf("Email body: %s", message)
	err := sendEmail(name, email, message)
	if err != nil {
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		log.Println("Error sending email:", err)
		return
	}

	http.Redirect(w, r, "/?status=success", http.StatusSeeOther)

}
func sendEmail(name, email, message string) error {
	log.Printf("Email body: %s", name)
	log.Printf("Email body: %s", email)
	log.Printf("Email body: %s", message)
	from := "jasso.favio@student.ccm.edu"
	password := "pwyaqdppgmdsewza"
	to := "jasso.favio@student.ccm.edu"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	body := fmt.Sprintf("Name: %s\nEmail: %s\nMessage: %s", name, email, message)

	auth := smtp.PlainAuth("", from, password, smtpHost)
	msg := []byte("To: " + to + "\r\n" +
		"Subject: Form Submission\r\n" +
		"\r\n" +
		body + "\r\n")

	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", formHandler)
	http.HandleFunc("/send", sendHandler)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
