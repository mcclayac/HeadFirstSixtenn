package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines
}

func viewHandler_old(writer http.ResponseWriter, request *http.Request) {
	placeholder := []byte("signature list goes here")
	_, err := writer.Write(placeholder)
	check(err)
}

type GuestBook struct {
	Signatures     []string
	SignatureCount int
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := getStrings("signatures.txt")
	fmt.Printf("%#v\n", signatures)

	guestBook := GuestBook{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}

	html, err := template.ParseFiles("view.html")
	check(err)

	err = html.Execute(writer, guestBook)
	check(err)
}

func newHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("new.html")
	check(err)

	err = html.Execute(writer, nil)
}

func viewHandler2(writer http.ResponseWriter, request *http.Request) {
	signatures := getStrings("signatures.txt")
	fmt.Printf("%#v\n", signatures)

	html, err := template.ParseFiles("view.html")
	check(err)

	err = html.Execute(writer, nil)
	check(err)
}

func createHandler(writer http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("signature")
	osOptions := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("signatures.txt", osOptions, os.FileMode(0600))
	check(err)

	_, err = fmt.Fprintln(file, signature)
	check(err)

	err = file.Close()
	check(err)
	// Redirect
	http.Redirect(writer, request, "/guestbook", http.StatusFound)

}

func guestBook1() {
	http.HandleFunc("/guestbook", viewHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

func guestBook2() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)

}

func main() {
	guestBook2()

}
