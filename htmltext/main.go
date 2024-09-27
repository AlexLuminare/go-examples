package main

import (
	"html/template"
	"os"
)

func main() {
	allTemplates, err1 := template.ParseFiles("templates/template.html",
		"templates/extras.html")
	if err1 == nil {
		allTemplates.ExecuteTemplate(os.Stdout,
			"template.html", &Kayak)
		os.Stdout.WriteString("\n")
		allTemplates.ExecuteTemplate(os.Stdout,
			"extras.html", &Kayak)
	} else {
		Printfln("Error: %v %v", err1.Error())
	}
}
