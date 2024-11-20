package domain

import (
	"mime/multipart"
	"time"
)

type Blog struct {
	ID        int                   `db:"id" form:"id"`
	Title     string                `db:"title" form:"title"`
	ImageLink string                `db:"image_link" form:"image_link"`
	Image     *multipart.FileHeader `form:"image_file"`
	Content   string                `db:"content" form:"content"`
	CreatedAt time.Time             `db:"created_at" form:"created_at"`
	UpdatedAt time.Time             `db:"updated_at" form:"updated_at"`
}
