package solid

import "fmt"

// Bad Code
type Report struct {
	Title string
	Data  string
}

func (r Report) SaveToFile() {
	fmt.Println("Saving report to file...")
}

func (r Report) SendEmail() {
	fmt.Println("Sending report through email...")
}

func (r Report) Generate() string {
	return fmt.Sprintf("Report: %s -> %s", r.Title, r.Data)
}

// What’s wrong?
// Report:
// generates content
// saves to file
// sends email
// It violates SRP because it has multiple reasons to change.

type FileSaver struct{}

func (FileSaver) SaveToFile(r Report) {
	fmt.Println("Saving report to file...")
}

type EmailSender struct{}

func (EmailSender) SendEmail(r Report) {
	fmt.Println("Sending report through email...")
}
