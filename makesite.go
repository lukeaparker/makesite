package main

import (
	  "fmt"
	  "bytes"
	  "io/ioutil"
	  "html/template"
	  "strings"
	  "os"
)

type Content struct {
	ParText string
  }

func main() {
		filename := os.Args[1]
        posText, err := ioutil.ReadFile(filename)
        if err != nil {
            panic(err)
		}
		data := Content{ParText: string(posText)}
		paths := []string{
			"template.tmpl",
		  }
		  buff := new(bytes.Buffer)
		  t := template.Must(template.New("template.tmpl").ParseFiles(paths...))
		  err = t.Execute(buff, data)		  
		  if err != nil {
		  	panic(err)
		  }
		  filename = strings.Replace(filename, ".txt", ".html", 1)
		  bytesToWrite := []byte(buff.Bytes())
		  err = ioutil.WriteFile(filename, bytesToWrite, 0644)
		  if err != nil {
			  panic(err)
		  }
		  fmt.Println(buff)

}