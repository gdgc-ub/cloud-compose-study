package server

import (
	"log"

	"github.com/Masterminds/sprig/v3"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"github.com/devanfer02/go-blog/internal/app/controller"
	"github.com/devanfer02/go-blog/internal/app/repository"
	"github.com/devanfer02/go-blog/internal/app/service"
	"github.com/devanfer02/go-blog/internal/infra/env"
)

type Server interface {
	MountMiddlewares()
	MountControllers()
	Start()
}

type httpServer struct {
	app *gin.Engine
	dbx *sqlx.DB
}

func NewHTTPServer(dbx *sqlx.DB) Server {
	app := gin.Default()

	return &httpServer{
		app: app,
		dbx: dbx,
	}
}

func (h *httpServer) MountMiddlewares() {
	h.app.SetFuncMap(sprig.FuncMap())
	h.app.Static("/static", "./resources/static")
	h.app.LoadHTMLGlob("resources/views/*")
}

func (h *httpServer) MountControllers() {
	// repositories
	blogRepo := repository.NewPgsqlBlogRepository(h.dbx)

	// services
	blogSvc := service.NewBlogService(blogRepo)

	// controllers
	controller.MountBlogRoutes(h.app, blogSvc)
}

func (h *httpServer) Start() {
	if env.AppEnv.AppPort[0] != ':' {
		env.AppEnv.AppPort = ":" + env.AppEnv.AppPort
	}

	if err := h.app.Run(env.AppEnv.AppPort); err != nil {
		log.Fatal("ERR: " + err.Error())
	}
}
