package main

import (
	"github.com/devanfer02/go-blog/infra/database"
	"github.com/devanfer02/go-blog/infra/server"
)

func main() {
	pgsqldb := database.NewPgsqlConn()
	httpSrv := server.NewHTTPServer(pgsqldb)

	httpSrv.MountMiddlewares()
	httpSrv.MountControllers()
	httpSrv.Start()
}
