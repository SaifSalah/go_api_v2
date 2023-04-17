package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/saifsalah/go_api_v2/internal/comment"
)

type CommentService interface {
	PostComment(context.Context, comment.Comment) (comment.Comment, error)
	GetComment(ctx context.Context, id string) (comment.Comment, error)
	//UpdateComment(ctx comment.Comment, id string, cmt comment.Comment) (comment.Comment, error)
	//DeleteComment(ctx context.Context, id string) error
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {

	var cmt comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		return
	}

	cmt, err := h.Service.PostComment(r.Context(), cmt)

	if err != nil {
		log.Print(err)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}

}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmt, err := h.Service.GetComment(r.Context(), id)

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return //STOP HERE :|
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}

}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {

}
