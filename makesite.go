package main

import ( 
	"fmt"
	"io/ioutil"
	"html/template"
	"bytes"
)
type Content struct {
	ParText string
	}

func main() {
	postText, err := ioutil.ReadFile("first-post.txt")
	content := Content{ParText: string(postText)}
	if err != nil {
		// A common use of `panic` is to abort if a function returns an error
		// value that we don’t know how to (or want to) handle. This example
		// panics if we get an unexpected error when creating a new file.
		panic(err)
		}
		paths := []string{
		"template.tmpl"}
		buff := new(bytes.Buffer)
		t := template.Must(template.New("template.tmpl").ParseFiles(paths...))
		err = t.Execute(buff, content)
		fmt.Print(buff.String())
		if err != nil {
			panic(err)
		}
		bytesToWrite := buff.Bytes()
        err = ioutil.WriteFile("new-file.txt", bytesToWrite, 0644)
        if err != nil {
            panic(err)
        }
}



