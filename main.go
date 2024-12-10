package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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

	mux.HandleFunc("/sendemail", SendEmail)

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

	type respT struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	resp := respT{
		Success: false,
		Message: "Inca nu!",
	}

	j, _ := json.Marshal(resp)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", string(j))

}
