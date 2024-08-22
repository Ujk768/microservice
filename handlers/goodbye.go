package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type GoodBye struct {
	l *log.Logger
}

func NewGoodBye(l *log.Logger) *GoodBye {
	return &GoodBye{l}
}

func (g *GoodBye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Print("Inside GoodBye")
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Something went wrong !!!", http.StatusBadRequest)
	}
	fmt.Fprintf(rw, "GoodBye %s", d)
}
