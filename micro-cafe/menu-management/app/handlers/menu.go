package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/amir-mln/go-projects/micro-cafe/menu-management/app"
)

const MenuHandlerRoute = "/menus"

type MenuHandler struct {
	l  *log.Logger
	as *app.ApplicationService
}

func NewMenuHandler(l *log.Logger, s *app.ApplicationService) *MenuHandler {
	return &MenuHandler{l, s}
}

func (h *MenuHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.handleGet(rw, r)
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *MenuHandler) handleGet(rw http.ResponseWriter, r *http.Request) {
	data, err := h.as.MenuBusiness.GetAllMenus()

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)

	encoder.SetIndent("", "  ")
	encoder.Encode(data)

	rw.Write([]byte(buffer.Bytes()))
}
