package restaurant

import (
	"encoding/json"
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
	router.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.get)
		r.Put("/", h.replace)
		r.Delete("/", h.remove)
	})

	return router
}

func (h *HttpHandler) RegisterTo(router chi.Router) {
	router.Route("/restaurants", func (r chi.Router) {
		r.Post("/", h.add)
		r.Get("/", h.getAll)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", h.get)
			r.Put("/", h.replace)
			r.Delete("/", h.remove)
			// r.Route("/menus", func(r chi.Router) {
			//  r.Post("/", h.add)
			// 	r.Get("/", h.getAll)
			// 	r.Get("/{id}", h.get)
			// 	r.Put("/{id}", h.replace)
			// 	r.Delete("/{id}", h.remove)
			// })
    })
	})
}

func (h *HttpHandler) add(w http.ResponseWriter, r *http.Request) {
	var restaurant Restaurant

	if err := json.NewDecoder(r.Body).Decode(&restaurant); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.repository.Add(r.Context(), &restaurant); err != nil {
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

func (h *HttpHandler) replace(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var restaurant Restaurant
	restaurant.ID = id
	if err := json.NewDecoder(r.Body).Decode(&restaurant); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.repository.Replace(r.Context(), &restaurant); err != nil {
		switch err {
		case ErrNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	render.Status(r, http.StatusAccepted)
	render.JSON(w, r, restaurant)
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

func (h *HttpHandler) get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	restaurant, err := h.repository.Get(r.Context(), id)
	if err != nil {
		switch err {
		case ErrNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, restaurant)
}

func (h *HttpHandler) getAll(w http.ResponseWriter, r *http.Request) {
	restaurants, err := h.repository.GetAll(r.Context())

	if err != nil {
		switch err {
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, restaurants)
}
