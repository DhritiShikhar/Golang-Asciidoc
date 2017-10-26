package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/VonC/asciidocgo"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func convert(w http.ResponseWriter, r *http.Request) {
	lines, err := readLines("random.adoc")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	_ = asciidocgo.NewDocument(lines, nil)

	//tmpl := template.Must(template.ParseFiles("layout.html"))

	rend := asciidocgo.Renderer{}
	for i := 0; i < len(lines); i++ {
		output := rend.Render("layout.html", lines[0], nil)
		io.WriteString(w, "_____________________________1________\n\n\n")
		io.WriteString(w, lines[i])
		io.WriteString(w, "_____________________________2________\n\n\n")
		io.WriteString(w, output)
	}
}

func main() {
	http.HandleFunc("/", convert)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
