package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrNotImplemented  = errors.New("not implemented")
)

type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

type Store interface {
	GetComment(ctx context.Context, uuid string) (Comment, error)
	PostComment(context.Context, Comment) (Comment, error)
}

type Service struct {
	Store Store
}

// NewService  is constructor
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retrieving comment")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}

	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, id string, comment Comment) error {
	return ErrNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string, cmt Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}

func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {

	insert, err := s.Store.PostComment(ctx, cmt)
	if err != nil {
		return Comment{}, fmt.Errorf("HERE %w", err)
	}
	fmt.Println(cmt)
	return insert, nil
}
