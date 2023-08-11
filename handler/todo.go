package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
	"github.com/TechBowl-japan/go-stations/service"
)

// A TODOHandler implements handling REST endpoints.
type TODOHandler struct {
	svc *service.TODOService
}

// NewTODOHandler returns TODOHandler based http.Handler.
func NewTODOHandler(svc *service.TODOService) *TODOHandler {
	return &TODOHandler{
		svc: svc,
	}
}

func (t *TODOHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		req := &model.CreateTODORequest{}
		dec := json.NewDecoder(r.Body)
		if err := dec.Decode(&req); err != nil {
			log.Println(err)
			return
		}
		if req.Subject == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		todo, err := t.svc.CreateTODO(r.Context(), req.Subject, req.Description)
		if err != nil {
			log.Println(err)
			return
		}
		res := &model.CreateTODOResponse{
			TODO: *todo,
		}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Println(err)
			return
		}
	} else if r.Method == http.MethodPut {
		req := &model.UpdateTODORequest{}
		dec := json.NewDecoder(r.Body)
		if err := dec.Decode(&req); err != nil {
			log.Println(err)
			return
		}
		if req.Subject == "" || req.ID == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		todo, err := t.svc.UpdateTODO(r.Context(), req.ID, req.Subject, req.Description)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		res := &model.UpdateTODOResponse{
			TODO: *todo,
		}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Println(err)
			return
		}
	}
}

// Create handles the endpoint that creates the TODO.
func (h *TODOHandler) Create(ctx context.Context, req *model.CreateTODORequest) (*model.CreateTODOResponse, error) {
	_, _ = h.svc.CreateTODO(ctx, "", "")
	return &model.CreateTODOResponse{}, nil
}

// Read handles the endpoint that reads the TODOs.
func (h *TODOHandler) Read(ctx context.Context, req *model.ReadTODORequest) (*model.ReadTODOResponse, error) {
	_, _ = h.svc.ReadTODO(ctx, 0, 0)
	return &model.ReadTODOResponse{}, nil
}

// Update handles the endpoint that updates the TODO.
func (h *TODOHandler) Update(ctx context.Context, req *model.UpdateTODORequest) (*model.UpdateTODOResponse, error) {
	_, _ = h.svc.UpdateTODO(ctx, 0, "", "")
	return &model.UpdateTODOResponse{}, nil
}

// Delete handles the endpoint that deletes the TODOs.
func (h *TODOHandler) Delete(ctx context.Context, req *model.DeleteTODORequest) (*model.DeleteTODOResponse, error) {
	_ = h.svc.DeleteTODO(ctx, nil)
	return &model.DeleteTODOResponse{}, nil
}
