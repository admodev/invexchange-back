package server

import (
	"admodev/invexchange/cmd/gincmd"
	"admodev/invexchange/container/routes"
)

func RunServer() {
	r := gincmd.R

	routes.InitializeRoutes()

	r.Run()
}
