package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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

	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch)
	<-ch
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)
	log.Println("Shutting down")
	return nil
}
