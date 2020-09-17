package main

import (
	"log"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func example1() {
	text := "Here's my template!\n\n"
	tmpl, err := template.New("test").Parse(text)
	check(err)

	err = tmpl.Execute(os.Stdout, nil)
}

func example2() {
	templateText := "Template Start\nAction: {{.}}\nTemplate end\n\n"
	tmpl, err := template.New("test").Parse(templateText)
	check(err)

	err = tmpl.Execute(os.Stdout, "ABC")
	check(err)

	err = tmpl.Execute(os.Stdout, 42)
	check(err)

	err = tmpl.Execute(os.Stdout, true)
	check(err)

	err = tmpl.Execute(os.Stdout, "Theze Nutz!")
	check(err)

}

func executeTemplate(text string, data interface{}) {

	tmpl, err := template.New("test").Parse(text)
	check(err)

	err = tmpl.Execute(os.Stdout, data)
	check(err)

}

func example3() {
	executeTemplate("Dot is: {{ . }}!\n", "ABC")
	executeTemplate("Dot is: {{ . }}!\n", 123.5)
	executeTemplate("Dot is: {{ . }}!\n", "Theze nutz!")

}

func example4() {

	executeTemplate("{{ if . }}Dot is true! {{ end }} finished\n", true)
	executeTemplate("{{ if . }}Dot is true! {{ end }} finished\n", false)
}

func example5() {
	templateText := "Before the loop {{ . }}\n{{ range . }} in the loop: {{.}}\n{{end}}After loop: {{ . }}\n\n\n"
	executeTemplate(templateText, []string{"do", "re", "mi", "fa", "so"})

	templateText = "Prices:\n{{ range . }} ${{.}}\n{{end}}"
	executeTemplate(templateText, []float64{1.25, 0.99, 27})

	executeTemplate(templateText, []float64{})
	executeTemplate(templateText, nil)
}

type Part struct {
	Name  string
	Count int
}

func example6() {
	templateText := "Name: {{ .Name }}\tCount: {{ .Count }}\n"
	executeTemplate(templateText, Part{Name: "Fuss", Count: 5})
	executeTemplate(templateText, Part{Name: "Cables", Count: 2})

}

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
}

func example7() {
	templateText := "Name: {{ .Name }}\t{{ if .Active}}Rate:{{ .Rate }}{{end}}\n"
	executeTemplate(templateText, Subscriber{Name: "Tony", Rate: 2.33, Active: true})
	executeTemplate(templateText, Subscriber{Name: "Maxine", Rate: 0.33, Active: true})
	executeTemplate(templateText, Subscriber{Name: "Kristin", Rate: 9.33, Active: false})
}

func main() {

	example2()
	example3()
	example4()
	example5()
	example6()
	example7()
}
