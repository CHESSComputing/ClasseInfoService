package main

import (
	srvConfig "github.com/CHESSComputing/golib/config"
	server "github.com/CHESSComputing/golib/server"
	"github.com/CHESSComputing/golib/services"
	"github.com/gin-gonic/gin"
)

// Verbose defines verbosity level
var Verbose int

// global variables
var _foxdenUser services.UserAttributes

// helper function to setup our router
func setupRouter() *gin.Engine {
	routes := []server.Route{
		{Method: "GET", Path: "/translate", Handler: GetHandler, Authorized: false},
	}
	r := server.Router(routes, nil, "static", srvConfig.Config.ClasseIdData.WebServer)
	return r
}

// Server defines our HTTP server
func Server() {
	// init Verbose
	Verbose = srvConfig.Config.ClasseIdData.WebServer.Verbose

	// make a choice of foxden user
	switch srvConfig.Config.ClasseIdData.FoxdenUser.User {
	case "Maglab":
		_foxdenUser = &services.MaglabUser{}
	case "CHESS":
		_foxdenUser = &services.CHESSUser{}
	default:
		_foxdenUser = &services.CHESSUser{}
	}
	_foxdenUser.Init()

	// setup web router and start the service
	r := setupRouter()
	webServer := srvConfig.Config.ClasseIdData.WebServer
	server.StartServer(r, webServer)
}
