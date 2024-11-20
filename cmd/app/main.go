package main

import (
	"github.com/devanfer02/go-blog/internal/infra/database"
	"github.com/devanfer02/go-blog/internal/infra/server"
)

func main() {
	mysqldb := database.NewMySQLConn()
	httpSrv := server.NewHTTPServer(mysqldb)

	httpSrv.MountMiddlewares()
	httpSrv.MountControllers()
	httpSrv.Start()
}
