package main

import (
	//"flag"
	"log"
	"net/http"
	//"os"
	"path/filepath"
	"sync"
	"text/template"

	//"github.com/seungkyua/hello-go/chat/trace"
)

type templateHandler struct {
	once		sync.Once
	filename	string
	templ		*template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	r := newRoom()

	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)

	r.run()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
