package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

// templ is a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("template", t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	r := newRoom()
	//endpoints
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	//get the room going
	go r.run()
	//start server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
