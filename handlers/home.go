package handlers

import (
	"html/template"
	"net/http"
	"time"

	"hestia/i18n"
	"hestia/static"
)

type PageData struct {
	Date string
	Text i18n.Translations
}

func Home() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		
		tmpl, err := template.ParseFiles(static.File("templates/index.html"))
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}
		
		translations, err := i18n.Load()
		if err != nil {
			http.Error(w, "Error loading translations", http.StatusInternalServerError)
			return
		}
		
		pageData := PageData{
			Date: time.Now().Format("2006-01-02"),
			Text: translations,
		}
		
		err = tmpl.Execute(w, pageData)
		if err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
			return
		}
	})
}