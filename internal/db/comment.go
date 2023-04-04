package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/saifsalah/go_api_v2/internal/comment"
	uuid "github.com/satori/go.uuid"
)

type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Body   sql.NullString
	Author sql.NullString
}

func convertCommentToRowComment(c CommentRow) comment.Comment {
	return comment.Comment{
		ID:     c.ID,
		Slug:   c.Slug.String,
		Author: c.Author.String,
		Body:   c.Body.String,
	}
}

func (d *Database) GetComment(ctx context.Context, uuid string) (comment.Comment, error) {
	var cmtRow CommentRow
	row := d.Client.QueryRowContext(
		ctx,
		`SELECT id,slug,body,author FROM comments WHERE id = $1`,
		uuid,
	)

	err := row.Scan(&cmtRow.ID, &cmtRow.Slug, &cmtRow.Body, &cmtRow.Author)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("Error featching the comment by uuid %w", err)
	}

	return convertCommentToRowComment(cmtRow), nil
}

func (d *Database) PostComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error) {

	cmt.ID = uuid.NewV4().String()
	postRow := CommentRow{
		ID:     cmt.ID,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
	}

	rows, err := d.Client.NamedQueryContext(ctx, `INSERT INTO comments (id,slug,author,body) VALUES (:id,:slug,:author,:body)`, postRow)

	if err != nil {
		return comment.Comment{}, fmt.Errorf("Fail to insert comment %w", err)
	}

	if err := rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("Fail to close rows: %w", err)
	}

	return cmt, nil
}
