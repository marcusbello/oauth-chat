package main

import (
	"flag"
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
	err := t.templ.Execute(w, r)
	if err != nil {
		return
	}
}

func main() {
	var addr = flag.String("addr", ":8080", "The addr ofthe application")
	flag.Parse()
	r := newRoom()
	//endpoints
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	//get the room going
	go r.run()
	//start server
	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
