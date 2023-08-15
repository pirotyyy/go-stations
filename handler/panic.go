package handler

import (
	"fmt"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
)

type PanicHandler struct{}

func NewPanicHandler() *PanicHandler {
	return &PanicHandler{}
}

func (h *PanicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	k := model.Key("os")
	fmt.Println(r.Context().Value(k))
	panic("Panic!")
}
