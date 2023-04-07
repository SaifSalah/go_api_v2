package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router  *mux.Router
	Service CommentService
	Server  *http.Server
}

type CommentService interface{}

func NewHandler(service CommentService) *Handler {

	handler := &Handler{
		Service: service,
	}
	handler.Router = mux.NewRouter()
	handler.mapRoutes()
	handler.Server = &http.Server{
		Addr:    "0.0.0.0:1992",
		Handler: handler.Router,
	}
	return handler
}

func (h *Handler) mapRoutes() {

	h.Router.HandleFunc("/helth", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "HELLo")
	})
}

func (h *Handler) Serve() error {

	if err := h.Server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
