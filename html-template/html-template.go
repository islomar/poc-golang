package main

import (
	"os"
	"html/template"
)

const html = `
<script>var foo = {{.Foo}};</script>  	\\double quotes are escaped
<a href="{{.URL}}">						\\URL gets encoded
	{{.Text}}   						\\the < becomes an HTML entity
</a>
`

func main() {
	tmpl := template.Must(template.New("example").Parse(html))
	data := struct {
		Foo			string
		URL, Text 	string
	}{
		Foo:	`Some "quoted" string`,  
		URL:	`" onClick="alert('xss!');`,  
		Text:	"The <- operator is for channel sends and receives",
	}
	tmpl.Execute(os.Stdout, data)
}