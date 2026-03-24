package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type HomeHandler struct {
	tmpl *template.Template
}

func NewHomeHandler(templatePath string) *HomeHandler {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatalf("error cargando template: %v", err)
	}
	return &HomeHandler{tmpl: tmpl}
}

func (h *HomeHandler) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := h.tmpl.Execute(w, nil); err != nil {
		http.Error(w, "error renderizando página", http.StatusInternalServerError)
	}
}
