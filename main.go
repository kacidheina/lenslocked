package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func homeHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my great site!</h1>")
}

func contactHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in contact with me, email me at <a "+
		" href = \"mailto:jon@calhoun.io\">jon@calhoun.io </a></p>")
}

func faqHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
<ul>
  <li>
    <b>Is there a free version?</b>
    Yes! We offer a free trial for 30 days on any paid plans.
  </li>
  <li>
    <b>What are your support hours?</b>
    We have support staff answering emails 24/7, though response
    times may be a bit slower on weekends.
  </li>
  <li>
    <b>How do I contact support?</b>
    Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>
  </li>
</ul>
`)
}

func galleryHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	id := chi.URLParam(req, "id")
	fmt.Fprintf(w, "<h1>Gallery Page</h1><p>ID: %s</p>", id)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", homeHandler)
	r.Get("/galleries/{id}", galleryHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000......")

	http.ListenAndServe(":3000", r)
}
