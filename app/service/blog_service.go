package service

import (
	"github.com/devanfer02/go-blog/app/repository"
	"github.com/devanfer02/go-blog/domain"
)

type BlogService interface {
	GetAllBlogs() ([]domain.Blog, error)
	GetBlogByID(id int) (domain.Blog, error)
	CreateBlog(blog *domain.Blog) error 
	UpdateBlog(blog *domain.Blog) error 
	DeleteBlog(id int) error 
}

type blogService struct {
	blogRepo repository.BlogRepository
}

func NewBlogService(blogRepo repository.BlogRepository) BlogService {
	return &blogService{blogRepo: blogRepo}
}

func(s *blogService) GetAllBlogs() ([]domain.Blog, error) {
	blogs, err := s.blogRepo.FetchAllBlogs()

	if err != nil {
		return nil, err 
	}

	return blogs, err 
}

func(s *blogService) GetBlogByID(id int) (domain.Blog, error) {
	blog, err := s.blogRepo.FetchBlogByID(id)
	
	if err != nil {
		return domain.Blog{}, err 
	}

	return blog, nil 
}

func(s *blogService) CreateBlog(blog *domain.Blog) error  {
	err := s.blogRepo.InsertBlog(blog)

	if err != nil {
		return err 
	}

	return nil 
}

func(s *blogService) UpdateBlog(blog *domain.Blog) error  {
	err := s.blogRepo.UpdateBlog(blog)

	if err != nil {
		return err 
	}

	return nil 
}

func(s *blogService) DeleteBlog(id int) error  {
	err := s.blogRepo.DeleteBlog(id)

	if err != nil {
		return err 
	}

	return nil 
}
