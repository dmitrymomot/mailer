package main

import (
	"log"
	"net/http"

	"github.com/dmitrymomot/mailer/template"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/assets/logo.svg", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./assets/logo.svg")
	})
	r.Get("/", preview)
	r.Get("/confirm-email", confirmEmailPreview)

	if err := http.ListenAndServe(":8080", r); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func preview(w http.ResponseWriter, r *http.Request) {
	err := template.ComponentsPreview().Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func confirmEmailPreview(w http.ResponseWriter, r *http.Request) {
	err := template.ConfirmEmail(template.ConfirmEmailPayload{
		Lang:         "en",
		Subject:      "Confirm your email",
		EmailPreview: "Click the button below to confirm your email address.",
		LogoURL:      "http://localhost:8080/assets/logo.svg",
		CompanyName:  "Acme Inc.",
		CompanyURL:   "http://localhost:8080/",
		Greeting:     "Hey there!",
		SupportEmail: "support@mail.dev", // fake email address for the example
		SupportURL:   "http://localhost:8080/",
		ButtonText:   "Confirm email",
		ButtonURL:    "http://localhost:8080/",
		ExpiresIn:    "24 hours",
		Footer:       "&copy; 2021 Acme Inc. All rights reserved. <br>1234 Acme St, Acmeville, AC 12345, USA <br>Phone: +123 456 789 000",
	}).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
