package main

import (
	"log"
	"net/http"
	"text/template"

	"os"

	"github.com/gorilla/sessions"
)

// session variable. (not used)
var cs *sessions.CookieStore = sessions.NewCookieStore([]byte("secret-key-1234"))

// Template for no-template.
func notemp() *template.Template {
	src := "<html><body><h1>NO TEMPLATE.</h1></body></html>"
	tmp, _ := template.New("home").Parse(src)
	return tmp
}

// get target Temlate.
func page(fname string) *template.Template {
	tmps, _ := template.ParseFiles("templates/"+fname+".html",
		"templates/head.html", "templates/foot.html")
	return tmps
}

// home handler.
func home(w http.ResponseWriter, rq *http.Request) {
	item := struct {
		Template string
		Title    string
		Message  string
	}{
		Template: "home",
		Title:    "Home",
		Message:  "出演者を入力",
	}
	er := page("home").Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
}

// tvlist handler.
func tvlist(w http.ResponseWriter, rq *http.Request) {
	data := []string{
		"6/15", "6/16", "6/17", "6/18", "6/19", "6/20", "6/21",
	}

	item := struct {
		Title string
		Data  []string
	}{
		Title: "番組表",
		Data:  data,
	}

	er := page("tvlist").Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
}

// main program.
func main() {
	dir, _ := os.Getwd()
	log.Print(http.Dir(dir + "/static/"))
	// home handling.
	http.HandleFunc("/", func(w http.ResponseWriter, rq *http.Request) {
		home(w, rq)
	})
	// tvlist handling
	http.HandleFunc("/tvlist", func(w http.ResponseWriter, rq *http.Request) {
		tvlist(w, rq)
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(dir+"/static/"))))

	http.ListenAndServe(":8080", nil)
}
