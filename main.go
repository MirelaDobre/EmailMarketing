package main

import (
	"fmt"
	"net/http"
	"net/smtp"

	"github.com/rs/cors"
)

func main() {

	fmt.Println("Start aplicatie!")

	_, err := DatabaseInitGC() //apelare functie initializare DB din fisierul database.go
	if err != nil {
		fmt.Println("Eroare la initializarea bazei de date." + err.Error())
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/sendemail/{tip}", SendEmail)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"Origin", "Content-Length", "Content-Type", "Cache-Control", "Authorization", "x-requested-with"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: false,
	})

	handler := c.Handler(mux)
	err = http.ListenAndServe(":80", handler)
	if err != nil {
		fmt.Println("Eroare la pornirea serverului." + err.Error())
	}

}

func SendEmail(w http.ResponseWriter, r *http.Request) {
	// Send email code here
	tip := r.PathValue("tip")

	q := "SELECT id, email, tip FROM email WHERE tip=?"
	if tip != "craciun" && tip != "oferte" {
		fmt.Fprintf(w, "%s", "Tip invalid!")
		return
	}

	rows, err := Db.Query(q, tip)

	if err != nil {
		fmt.Fprintf(w, "%s", "Eroare la interogare baza de date!")
		return
	}

	defer rows.Close()
	for rows.Next() {
		var id int
		var email, t string
		err = rows.Scan(&id, &email, &t)
		if err != nil {
			fmt.Fprintf(w, "%s", "Eroare la citirea datelor!"+err.Error())
			return
		}

		smtpServer := "smtp.gmail.com"
		port := "587"
		username := "alexm.dobre@gmail.com"
		password := "atfq nkbb jfoo zhnd"

		// Email headers and body
		headers := "From: " + username + "\r\n" +
			"To: " + email + "\r\n" +
			"Subject: Email Marketing \r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n"
		message := headers + GetEmail(tip)

		// Send email
		auth := smtp.PlainAuth("", username, password, smtpServer)
		err := smtp.SendMail(smtpServer+":"+port, auth, username, []string{email}, []byte(message))
		if err != nil {
			fmt.Fprintf(w, "%s", "Eroare la trimitere email!"+err.Error())
		} else {
			fmt.Fprintf(w, "%s%s%s%s%s\r\n", "Email catre ", email, " de tip ", tip, " trimis cu succes!")
		}

		//fmt.Println(id, email, tip)
	}

	fmt.Fprintf(w, "%s", "Email-uri trimise cu succes!")

}
