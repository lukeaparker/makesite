package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
)

type Content struct {
	ParText string
}

func main() {
	buff := new(bytes.Buffer)
	dir := flag.String("dir", ".", "Name of the directory to save the File")
	fileName := flag.String("file", "first-post.txt", "name of file to write to html")

	flag.Parse()

	files, err := ioutil.ReadDir(*dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".txt" {
			fmt.Println(file.Name())

			posText, err := ioutil.ReadFile(file.Name())
			if err != nil {
				panic(err)
			}
			data := Content{ParText: string(posText)}
			paths := []string{
				"template.tmpl",
			}
			t := template.Must(template.New("template.tmpl").ParseFiles(paths...))
			err = t.Execute(buff, data)
			if err != nil {
				panic(err)
			}
			filename := strings.Replace(file.Name(), ".txt", ".html", 1)
			bytesToWrite := []byte(buff.Bytes())
			err = ioutil.WriteFile(filename, bytesToWrite, 0644)
			if err != nil {
				panic(err)
			}
		}
		for _, file := range files {
			if filepath.Ext(file.Name()) == ".md" {
				fileContents, err := ioutil.ReadFile(*fileName)
				if err != nil {
					panic(err)
				}
				mdBytes := []byte(fileContents)
				output := markdown.ToHTML(mdBytes, nil, nil)
				fmt.Println(output)
				buff.Reset()
			}
		}
		fmt.Println(buff)
		fmt.Println(fileName)

	}
}
