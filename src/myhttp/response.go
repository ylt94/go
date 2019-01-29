package myhttp

import(
	"html/template"
	"net/http"
)

func View(response http.ResponseWriter,path string,data []string){
	t, _ := template.ParseFiles(path)
    t.Execute(response, data)
}