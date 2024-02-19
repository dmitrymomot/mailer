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
	r.Get("/assets/logo.svg", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./assets/logo.svg")
	})
	r.Get("/", uiPreview)
	r.Get("/confirm", confirmEmailPreview)
	r.Get("/reset", resetPasswordPreview)

	if err := http.ListenAndServe(":8080", r); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func confirmEmailPreview(w http.ResponseWriter, r *http.Request) {
	err := mailer_template.Main(
		mailer_template.MainData{
			Title:     "Confirm Email Preview",
			Preheader: "Confirm your email address",
		},
		mailer_template.Logo("http://localhost:8080/assets/logo.svg", "Example", "180px", "http://localhost:8080"),
		mailer_template.Card(
			mailer_template.SubHeader("Dear [Recipient's Name],"),
			mailer_template.Paragraph("Thank you for [registering for our event/signing up for our service/another action]! We're excited to have you on board. To complete the process, please confirm your registration by clicking the button below:"),
			mailer_template.ButtonPrimary(mailer_template.ButtonAlignCenter, "Confirm email", "https://example.com"),
			mailer_template.Paragraph("Once confirmed, you will gain full access to [specific benefits, features, services, etc.]. Here's a quick overview of what to expect:"),
			mailer_template.List(
				"Lorem ipsum dolor sit amet",
				"Consectetur adipiscing elit",
				"Sed do eiusmod tempor incididunt",
				"Ut labore et dolore magna aliqua",
			),
			mailer_template.Paragraph("If you have any questions or need further assistance, please do not hesitate to contact us at [Contact Email] or visit our [FAQ Page/Help Center]."),
			mailer_template.Paragraph("Thank you for choosing Company. We look forward to providing you with the best experience."),
			mailer_template.Paragraph("Best regards, Company Name Team"),
			mailer_template.Divider(),
			mailer_template.ParagraphSecondary("If the button above does not work, please copy and paste the following link into your browser: "),
			mailer_template.Link("https://example.com/dfada/adawfawef/adsgafae?sfawefa=2342&asdasda=aefaedwfeggntyhrt.afafaegfwaefewf.shertgergrgargar&test=1", ""),
		),
		mailer_template.ParagraphSecondary("@ 2024 All rights reserved"),
	).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func resetPasswordPreview(w http.ResponseWriter, r *http.Request) {}

func uiPreview(w http.ResponseWriter, r *http.Request) {
	err := mailer_template.Main(
		mailer_template.MainData{
			Title:     "Test Email Preview",
			Preheader: "Enim deserunt non deserunt deserunt commodo cupidatat excepteur esse ad laborum.",
		},
		mailer_template.Container(
			mailer_template.AlertInfo("Aliqua adipisicing enim irure do consectetur nostrud exercitation consectetur commodo."),
		),
		mailer_template.Container(
			mailer_template.AlertInfoOutlined("Dolore exercitation voluptate cupidatat cillum quis exercitation irure adipisicing ex."),
		),
		mailer_template.Container(

			mailer_template.AlertSuccess("Aliqua adipisicing enim irure do consectetur nostrud exercitation consectetur commodo."),
		),
		mailer_template.Container(
			mailer_template.AlertSuccessOutlined("Dolore exercitation voluptate cupidatat cillum quis exercitation irure adipisicing ex."),
		),
		mailer_template.Container(
			mailer_template.AlertWarning("Aliqua adipisicing enim irure do consectetur nostrud exercitation consectetur commodo."),
		),
		mailer_template.Container(
			mailer_template.AlertWarningOutlined("Dolore exercitation voluptate cupidatat cillum quis exercitation irure adipisicing ex."),
		),
		mailer_template.Container(
			mailer_template.AlertDanger("Aliqua adipisicing enim irure do consectetur nostrud exercitation consectetur commodo."),
		),
		mailer_template.Container(
			mailer_template.AlertDangerOutlined("Dolore exercitation voluptate cupidatat cillum quis exercitation irure adipisicing ex."),
		),
		mailer_template.Logo("http://localhost:8080/assets/logo.svg", "Example", "200px", "http://localhost:8080"),
		mailer_template.Card(
			mailer_template.SubHeader("Hi there!"),
			mailer_template.Paragraph("Aute do aliqua consequat nulla ad adipisicing aliquip adipisicing exercitation esse. Ipsum amet consequat aliqua culpa excepteur sint anim dolor duis magna et dolor. Dolore Lorem ut dolor aliquip amet sint. Nostrud deserunt sunt elit amet qui consequat sint sit. Nostrud laborum dolore occaecat officia ex culpa officia cillum occaecat proident amet laboris culpa anim."),
			mailer_template.ButtonPrimary(mailer_template.ButtonAlignCenter, "Confirm email", "https://example.com"),
			mailer_template.Paragraph("Nulla excepteur ex eiusmod ad cupidatat laborum commodo consectetur nulla. Eu anim consectetur esse cupidatat irure irure. Anim aliqua irure eiusmod dolor Lorem non. Ea reprehenderit ex sunt in nulla fugiat amet ea <a href=\"http://example.com\">aliquip aliqua ad</a>."),
			mailer_template.List(
				"Lorem ipsum dolor sit amet",
				"Consectetur adipiscing elit",
				"Sed do eiusmod tempor incididunt",
				"Ut labore et dolore magna aliqua",
			),
			mailer_template.Paragraph("Reprehenderit ea excepteur cillum aute aliqua. Incididunt ex do anim ex fugiat eiusmod incididunt est culpa officia et dolore. Qui ipsum voluptate sunt nisi irure eiusmod incididunt dolore elit enim anim."),
			mailer_template.Divider(),
			mailer_template.ParagraphSecondary("Aute do aliqua consequat nulla ad adipisicing aliquip adipisicing exercitation esse. Ipsum amet consequat aliqua culpa excepteur sint anim dolor duis magna et dolor. Dolore Lorem ut dolor aliquip amet sint."),
		),
		mailer_template.Card(
			mailer_template.Header("Test Email Preview"),
			mailer_template.Paragraph("Aute do aliqua consequat nulla ad adipisicing aliquip adipisicing exercitation esse. Ipsum amet consequat aliqua culpa excepteur sint anim dolor duis magna et dolor. Dolore Lorem ut dolor aliquip amet sint. Nostrud deserunt sunt elit amet qui consequat sint sit. Nostrud laborum dolore occaecat officia ex culpa officia cillum occaecat proident amet laboris culpa anim."),
			mailer_template.ButtonDanger(mailer_template.ButtonAlignCenter, "Click me", "https://example.com"),
			mailer_template.Divider(),
			mailer_template.ButtonOutlinedDanger(mailer_template.ButtonAlignCenter, "Click me", "https://example.com"),
			mailer_template.CouponDashed(mailer_template.Coupon{
				Persentage:     10,
				Code:           "TESTCODE",
				ActionBtnURL:   "https://example.com",
				ActionBtnLabel: "Use code",
			}),
			mailer_template.Divider(),
			mailer_template.ParagraphSecondary("Aute do aliqua consequat nulla ad adipisicing aliquip adipisicing exercitation esse. Ipsum amet consequat aliqua culpa excepteur sint anim dolor duis magna et dolor. Dolore Lorem ut dolor aliquip amet sint. Nostrud deserunt sunt elit amet qui consequat sint sit. Nostrud laborum dolore occaecat officia ex culpa officia cillum occaecat proident amet laboris culpa anim."),
		),
		mailer_template.Container(
			mailer_template.Divider(),

			mailer_template.ParagraphSecondary("Aute do aliqua consequat nulla ad adipisicing aliquip adipisicing exercitation esse. Ipsum amet consequat aliqua culpa excepteur sint anim dolor duis magna et dolor. Dolore Lorem ut dolor aliquip amet sint. Nostrud deserunt sunt elit amet qui consequat sint sit. Nostrud laborum dolore occaecat officia ex culpa officia cillum occaecat proident amet laboris culpa anim."),
		),
		mailer_template.ParagraphSecondary("@ 2024 All rights reserved"),
	).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
