package repository

import (
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/devanfer02/go-blog/domain"
)

type BlogRepository interface {
	FetchAllBlogs() ([]domain.Blog, error)
	FetchBlogByID(id int) (domain.Blog, error)
	InsertBlog(blog *domain.Blog) error
	UpdateBlog(blog *domain.Blog) error
	DeleteBlog(id int) error
}

type pgsqlBlogRepository struct {
	conn *sqlx.DB
}

func NewPgsqlBlogRepository(conn *sqlx.DB) BlogRepository {
	return &pgsqlBlogRepository{conn}
}

const TABLE_NAME = "blogs"

func (r *pgsqlBlogRepository) FetchAllBlogs() ([]domain.Blog, error) {
	var (
		query sq.SelectBuilder
		sql   string
		err   error
		blogs []domain.Blog = make([]domain.Blog, 0)
	)

	query = sq.Select("*").From(TABLE_NAME)

	sql, _, err = query.ToSql()

	if err != nil {
		log.Printf("[BLOG REPOSITORY][FetchAllBlogs] ERR: %v\n", err.Error())
		return nil, err 
	}

	if err = r.conn.Select(&blogs, sql); err != nil {
		log.Printf("[BLOG REPOSITORY][FetchAllBlogs] ERR: %v\n", err.Error())
		return nil, err 
	}

	return blogs, nil 
}

func (r *pgsqlBlogRepository) FetchBlogByID(id int) (domain.Blog, error) {
	var (
		query sq.SelectBuilder
		sql string 
		err error 
		blogs []domain.Blog
		args []interface{}
	)

	query = sq.Select("*").From(TABLE_NAME).Where("id = ?", id)

	sql, args, err = query.PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		log.Printf("[BLOG REPOSITORY][FetchBlogsByID] ERR: %v\n", err.Error())
		return domain.Blog{}, err 
	}

	if err = r.conn.Select(&blogs, sql, args...); err != nil {
		log.Printf("[BLOG REPOSITORY][FetchBlogsByID] ERR: %v\n", err.Error())
		return domain.Blog{}, err 
	}

	if len(blogs) == 0 {
		return domain.Blog{}, domain.ErrNotFound
	}

	return blogs[0], nil 
}

func (r *pgsqlBlogRepository) InsertBlog(blog *domain.Blog) error {
	var (
		query sq.InsertBuilder
		sql string 
		err error
		args []interface{}
	)

	query = sq.
		Insert(TABLE_NAME). 
		Columns("title", "image_link", "content"). 
		Values(blog.Title, blog.ImageLink, blog.Content)

	sql, args, err = query.PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		log.Printf("[BLOG REPOSITORY][InsertBlog] ERR: %v\n", err.Error())
		return err 
	}

	if _, err = r.conn.Exec(sql, args...); err != nil {
		log.Printf("[BLOG REPOSITORY][InsertBlog] ERR: %v\n", err.Error())
		return err 
	}

	return nil
}

func (r *pgsqlBlogRepository) UpdateBlog(blog *domain.Blog) error {
	var (
		query sq.UpdateBuilder
		sql string 
		err error
		args []interface{}
	)

	query = sq.
		Update(TABLE_NAME). 
		Set("title", blog.Title). 
		Set("image_link", blog.ImageLink).
		Set("content", blog.Content). 
		Where("id = ?", blog.ID)

	sql, args, err = query.PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		log.Printf("[BLOG REPOSITORY][UpdateBlog] ERR: %v\n", err.Error())
		return err 
	}

	res, err := r.conn.Exec(sql, args...); 
	
	if err != nil {
		log.Printf("[BLOG REPOSITORY][DeleteBlog] ERR: %v\n", err.Error())
		return err 
	}

	if rows, _ := res.RowsAffected(); rows < 1 {
		return domain.ErrNotFound
	}

	return nil
}

func (r *pgsqlBlogRepository) DeleteBlog(id int) error {
	var (
		query sq.DeleteBuilder
		sql string 
		err error
		args []interface{}
	)

	query = sq.
		Delete(TABLE_NAME). 
		Where("id = ?", id)

	sql, args, err = query.PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		log.Printf("[BLOG REPOSITORY][DeleteBlog] ERR: %v\n", err.Error())
		return err 
	}

	res, err := r.conn.Exec(sql, args...); 
	
	if err != nil {
		log.Printf("[BLOG REPOSITORY][DeleteBlog] ERR: %v\n", err.Error())
		return err 
	}

	if rows, _ := res.RowsAffected(); rows < 1 {
		return domain.ErrNotFound
	}

	return nil
}
