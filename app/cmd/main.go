package main

import (
	"github.com/devanfer02/go-blog/infra/database"
	"github.com/devanfer02/go-blog/infra/server"
)

func main() {
	mysqldb := database.NewMySQLConn()
	httpSrv := server.NewHTTPServer(mysqldb)

	httpSrv.MountMiddlewares()
	httpSrv.MountControllers()
	httpSrv.Start()
}
