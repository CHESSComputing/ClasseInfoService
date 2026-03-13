package main

import (
	srvConfig "github.com/CHESSComputing/golib/config"
	ldap "github.com/CHESSComputing/golib/ldap"
	server "github.com/CHESSComputing/golib/server"
	"github.com/CHESSComputing/golib/services"
	"github.com/gin-gonic/gin"
)

// Verbose defines verbosity level
var Verbose int

// keep ldap cache
var ldapCache *ldap.Cache

// global variables
var _foxdenUser services.UserAttributes

// helper function to setup our router
func setupRouter() *gin.Engine {
	routes := []server.Route{
		{Method: "GET", Path: "/translate", Handler: GetHandler, Authorized: false},
	}
	r := server.Router(routes, nil, "static", srvConfig.Config.ClasseInfoData.WebServer)
	return r
}

// Server defines our HTTP server
func Server() {
	// init Verbose
	Verbose = srvConfig.Config.ClasseInfoData.WebServer.Verbose

	// initialize ldap cache
	ldapCache = &ldap.Cache{Map: make(map[string]ldap.Entry)}
	// make a choice of foxden user
	switch srvConfig.Config.ClasseInfoData.FoxdenUser.User {
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
	webServer := srvConfig.Config.ClasseInfoData.WebServer
	server.StartServer(r, webServer)
}
