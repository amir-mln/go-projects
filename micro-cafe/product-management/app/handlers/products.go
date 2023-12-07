package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/amir-mln/go-projects/micro-cafe/product-management/app"
)

const ProductHandlerRoute = "/products"

type ProductHandler struct {
	l *log.Logger
	s *app.ApplicationService
}

func NewProductHandler(l *log.Logger, s *app.ApplicationService) *ProductHandler {
	return &ProductHandler{l, s}
}

func (h *ProductHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.handleGet(rw, r)
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *ProductHandler) handleGet(rw http.ResponseWriter, r *http.Request) {
	data, err := h.s.PBiz.GetAllProducts()

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	HandlerBuffer := new(bytes.Buffer)
	json.NewEncoder(HandlerBuffer).Encode(data)

	rw.Write([]byte(HandlerBuffer.Bytes()))
}
