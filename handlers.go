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
	gidNumber := c.DefaultQuery("gidNumber", "")
	email := c.DefaultQuery("email", "")
	if name == "" && uid == "" && uidNumber == "" && email == "" && gidNumber == "" {
		rec := services.Response("ClasseInfoService",
			http.StatusBadRequest,
			services.LDAPSearchError,
			errors.New("no name or uid or uidNumber or email parameter is provided"))
		c.JSON(http.StatusBadRequest, rec)
		return
	}

	// make ldap query
	var err error
	var record ldap.Entry
	if uid != "" {
		record, err = ldapCache.SearchBy(
			srvConfig.Config.LDAP.Login,
			srvConfig.Config.LDAP.Password,
			uid, "uid")
		records = append(records, record)
	} else if name != "" {
		record, err = ldapCache.SearchBy(
			srvConfig.Config.LDAP.Login,
			srvConfig.Config.LDAP.Password,
			name, "name")
		records = append(records, record)
	} else if uidNumber != "" {
		record, err = ldapCache.SearchBy(
			srvConfig.Config.LDAP.Login,
			srvConfig.Config.LDAP.Password,
			uidNumber, "uidNumber")
		records = append(records, record)
	} else if gidNumber != "" {
		records, err = ldap.Records(
			srvConfig.Config.LDAP.Login,
			srvConfig.Config.LDAP.Password,
			gidNumber, "gidNumber", 0)
	} else if email != "" {
		records, err = ldap.Records(
			srvConfig.Config.LDAP.Login,
			srvConfig.Config.LDAP.Password,
			email, "mail", 0)
	}
	if err != nil {
		msg := fmt.Sprintf("No LDAP entry, error: %v", err)
		rec := services.Response("ClasseInfoService", http.StatusBadRequest, services.LDAPSearchError, errors.New(msg))
		c.JSON(http.StatusBadRequest, rec)
		return
	}
	c.JSON(http.StatusOK, records)
}
