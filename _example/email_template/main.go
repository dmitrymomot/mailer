package main

import (
	"log"
	"net/http"

	mailer_template "github.com/dmitrymomot/mailer/template"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", emailPreview)

	if err := http.ListenAndServe(":8080", r); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func emailPreview(w http.ResponseWriter, r *http.Request) {
	err := mailer_template.Main(
		mailer_template.MainData{
			Title:     "Test Email Preview",
			Preheader: "This is a test email preview",
		},
		mailer_template.Container(
			mailer_template.Alert("This is a test email preview"),
		),
		mailer_template.Container(
			mailer_template.AlertOutlined("This is a test email preview"),
		),
		mailer_template.Card(
			mailer_template.Header("Test Email Preview"),
			mailer_template.Paragraph("Aute do aliqua consequat nulla ad adipisicing aliquip adipisicing exercitation esse. Ipsum amet consequat aliqua culpa excepteur sint anim dolor duis magna et dolor. Dolore Lorem ut dolor aliquip amet sint. Nostrud deserunt sunt elit amet qui consequat sint sit. Nostrud laborum dolore occaecat officia ex culpa officia cillum occaecat proident amet laboris culpa anim."),
			mailer_template.ButtonPrimary(mailer_template.ButtonAlignCenter, "Click me", "https://example.com"),
			mailer_template.Divider(),
			mailer_template.ButtonOutlinedPrimary(mailer_template.ButtonAlignCenter, "Click me", "https://example.com"),
			mailer_template.CouponDashed(mailer_template.Coupon{
				Persentage:     10,
				Code:           "TESTCODE",
				ActionBtnURL:   "https://example.com",
				ActionBtnLabel: "Use code",
			}),
			mailer_template.Divider(),
			mailer_template.ParagraphSecondary("Aute do aliqua consequat nulla ad adipisicing aliquip adipisicing exercitation esse. Ipsum amet consequat aliqua culpa excepteur sint anim dolor duis magna et dolor. Dolore Lorem ut dolor aliquip amet sint. Nostrud deserunt sunt elit amet qui consequat sint sit. Nostrud laborum dolore occaecat officia ex culpa officia cillum occaecat proident amet laboris culpa anim."),
		),
		mailer_template.ParagraphSecondary("@ 2024 All rights reserved"),
	).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
