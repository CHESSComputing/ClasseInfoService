package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	srvConfig "github.com/CHESSComputing/golib/config"
	"github.com/CHESSComputing/golib/ldap"
	services "github.com/CHESSComputing/golib/services"
	"github.com/gin-gonic/gin"
)

// GetHandler handles queries via GET requests
func GetHandler(c *gin.Context) {
	var records []ldap.Entry

	name := c.DefaultQuery("name", "")
	uid := c.DefaultQuery("uid", "")
	if name == "" && uid == "" {
		rec := services.Response("ClasseInfoService",
			http.StatusBadRequest,
			services.LDAPSearchError,
			errors.New("no input name or uid parameter"))
		c.JSON(http.StatusBadRequest, rec)
		return
	}

	// make ldap query
	entry, err := ldapCache.Search(
		srvConfig.Config.LDAP.Login,
		srvConfig.Config.LDAP.Password,
		uid)
	log.Printf("ldap entry=%+v, err=%v", entry, err)
	if err != nil {
		msg := fmt.Sprintf("No LDAP entry, error: %v", err)
		rec := services.Response("ClasseInfoService", http.StatusBadRequest, services.LDAPSearchError, errors.New(msg))
		c.JSON(http.StatusBadRequest, rec)
		return
	}
	records = append(records, entry)

	c.JSON(http.StatusOK, records)
}
