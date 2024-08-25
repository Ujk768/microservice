package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

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
	if rq.Method == http.MethodPost {
		p.addProduct(rw, rq)
		return
	}
	//expect id in the URI
	if rq.Method == http.MethodPut {
		r := regexp.MustCompile(`/([0-9]+)`)
		g := r.FindAllStringSubmatch(rq.URL.Path, -1)
		if len(g) != 1 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		p.l.Println("Id:%v", id)
		p.updateProduct(id, rw, rq)
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

func (p *Products) addProduct(rw http.ResponseWriter, rq *http.Request) {
	p.l.Println("Iniside add Product hANDLER")
	prod := &data.Product{}
	err := prod.FromJSON(rq.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal JSON", http.StatusBadRequest)
	}
	data.AddProduct(*prod)
	fmt.Fprintf(rw, "Sucessfully inserted data")
}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, rq *http.Request) {
	p.l.Println("Iniside Update hANDLER")
	prod := &data.Product{}
	err := prod.FromJSON(rq.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal JSON", http.StatusBadRequest)
	}
	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNOtFound {
		http.Error(rw, "Product NOt found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product NOt found", http.StatusInternalServerError)
		return
	}
}
