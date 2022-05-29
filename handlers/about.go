package handlers

import (
	"log"
	"net/http"
)

type About struct {
	l *log.Logger
}

func NewAbout(l *log.Logger) *About {
	return &About{l}
}

func (a *About) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.l.Println("This is the about page")
	_, err := w.Write([]byte("This is the about page"))
	if err != nil {
		return
	}
}
