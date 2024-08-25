package handlers

import (
	"log"
	"net/http"

	"github.com/Ujk768/products/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// add logic to to which handler to move the request to
func (p *Products) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	if rq.Method == http.MethodGet {
		p.getProducts(rw, rq)
		return
	}
	
	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Products) getProducts(rw http.ResponseWriter, rq *http.Request) {
	lp := data.GetProducts()
	//convert list of products to JSON
	// d, err := json.Marshal(lp)
	// if err != nil {
	// 	http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	// }
	// rw.Write(d)
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}

}
