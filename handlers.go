package main

import (
	"errors"
	"fmt"
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
	uidNumber := c.DefaultQuery("uidNumber", "")
	email := c.DefaultQuery("email", "")
	if name == "" && uid == "" && uidNumber == "" && email == "" {
		rec := services.Response("ClasseInfoService",
			http.StatusBadRequest,
			services.LDAPSearchError,
			errors.New("no name or uid or uidNumber or email parameter is provided"))
		c.JSON(http.StatusBadRequest, rec)
		return
	}

	// make ldap query
	var entry ldap.Entry
	var err error
	if uid != "" {
		entry, err = ldapCache.SearchBy(
			srvConfig.Config.LDAP.Login,
			srvConfig.Config.LDAP.Password,
			uid, "uid")
	} else if name != "" {
		entry, err = ldapCache.SearchBy(
			srvConfig.Config.LDAP.Login,
			srvConfig.Config.LDAP.Password,
			name, "name")
	} else if uidNumber != "" {
		entry, err = ldapCache.SearchBy(
			srvConfig.Config.LDAP.Login,
			srvConfig.Config.LDAP.Password,
			uidNumber, "uidNumber")
	} else if email != "" {
		entry, err = ldapCache.SearchBy(
			srvConfig.Config.LDAP.Login,
			srvConfig.Config.LDAP.Password,
			email, "mail")
	}
	if err != nil {
		msg := fmt.Sprintf("No LDAP entry, error: %v", err)
		rec := services.Response("ClasseInfoService", http.StatusBadRequest, services.LDAPSearchError, errors.New(msg))
		c.JSON(http.StatusBadRequest, rec)
		return
	}
	records = append(records, entry)

	c.JSON(http.StatusOK, records)
}
