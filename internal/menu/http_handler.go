package menu

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type HttpHandler struct {
	repository Repository
}

func NewHttpHandler(repository Repository) *HttpHandler {
  return &HttpHandler{
		repository: repository,
	}
}

func (h *HttpHandler) Routes() chi.Router {
	router := chi.NewRouter()
	
	router.Post("/", h.add)
	router.Get("/", h.getAll)
	router.Get("/{id}", h.get)
	router.Delete("/{id}", h.remove)

	return router
}

func (h *HttpHandler) SubRoutes() chi.Router {
	router := chi.NewRouter()
	
	router.Get("/", h.getAllByRestaurant)
	router.Get("/{id}", h.get)

	return router
}

func (h *HttpHandler) add(w http.ResponseWriter, r *http.Request) {
	var menu Menu

	if err := json.NewDecoder(r.Body).Decode(&menu); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.repository.Add(r.Context(), &menu); err != nil {
		switch err {
		case ErrDuplicateIdentifier:
			http.Error(w, err.Error(), http.StatusConflict)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *HttpHandler) get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	menu, err := h.repository.Get(r.Context(), id)

	if err != nil {
		switch err {
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, menu)
}

func (h *HttpHandler) getAll(w http.ResponseWriter, r *http.Request) {
	menus, err := h.repository.GetAll(r.Context())

	if err != nil {
		switch err {
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, menus)
}

func (h *HttpHandler) getAllByRestaurant(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Printf(id)
	menus, err := h.repository.GetAllByRestaurant(r.Context(), id)

	if err != nil {
		switch err {
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, menus)
}

func (h *HttpHandler) remove(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := h.repository.Remove(r.Context(), id); err != nil {
		switch err {
		case ErrNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	render.NoContent(w, r)
}
