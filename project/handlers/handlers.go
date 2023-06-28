package handlers

import (
	"go-microservices/data"
	"net/http"
	"time"

	"github.com/wtran29/fenix/fenix"
)

type Handlers struct {
	App    *fenix.Fenix
	Models data.Models
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	defer h.App.LoadTime(time.Now())
	err := h.render(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}
