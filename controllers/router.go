package controllers

import (
	"testserver/models"

	"github.com/gin-gonic/gin"
	"github.com/littletwolee/commons"
)

var (
	Router *router
)

type router struct {
	logger *commons.Log
	Router *models.Router
}

func init() {
	if Router == nil {
		Router = &router{}
	}
	Router.logger = commons.GetLogger()
}
func (r *router) Post(c *gin.Context) {
	// var (
	// 	err  error
	// 	user models.User
	// )
	// err = c.Bind(&user)
	// if err != nil {
	// 	goto PARASERR
	// }
	// err = modules.User.Insert(&user)
	// if err != nil {
	// 	goto ERR
	// }
	c.JSON(200, c.Request.RequestURI)
	// 	goto RETURN
	// ERR:
	// 	r.logger.LogErr(nil)
	// 	c.JSON(commons.StatusInternalServerError, nil)
	// 	goto RETURN
	// PARASERR:
	// 	c.JSON(commons.StatusBadRequest, nil)
	goto RETURN
RETURN:
	return
}

func (r *router) Get(c *gin.Context) {
	// var (
	// 	err  error
	// 	user models.User
	// )
	// err = c.Bind(&user)
	// if err != nil {
	// 	goto PARASERR
	// }
	// err = modules.User.Insert(&user)
	// if err != nil {
	// 	goto ERR
	// }
	c.JSON(200, "get")
	// 	goto RETURN
	// ERR:
	// 	r.logger.LogErr(nil)
	// 	c.JSON(commons.StatusInternalServerError, nil)
	// 	goto RETURN
	// PARASERR:
	// 	c.JSON(commons.StatusBadRequest, nil)
	goto RETURN
RETURN:
	return
}

func (r *router) Put(c *gin.Context) {
	// var (
	// 	err  error
	// 	user models.User
	// )
	// err = c.Bind(&user)
	// if err != nil {
	// 	goto PARASERR
	// }
	// err = modules.User.Insert(&user)
	// if err != nil {
	// 	goto ERR
	// }
	c.JSON(200, "put")
	// 	goto RETURN
	// ERR:
	// 	r.logger.LogErr(nil)
	// 	c.JSON(commons.StatusInternalServerError, nil)
	// 	goto RETURN
	// PARASERR:
	// 	c.JSON(commons.StatusBadRequest, nil)
	goto RETURN
RETURN:
	return
}

func (r *router) Delete(c *gin.Context) {
	// var (
	// 	err  error
	// 	user models.User
	// )
	// err = c.Bind(&user)
	// if err != nil {
	// 	goto PARASERR
	// }
	// err = modules.User.Insert(&user)
	// if err != nil {
	// 	goto ERR
	// }
	c.JSON(200, "delete")
	// 	goto RETURN
	// ERR:
	// 	r.logger.LogErr(nil)
	// 	c.JSON(commons.StatusInternalServerError, nil)
	// 	goto RETURN
	// PARASERR:
	// 	c.JSON(commons.StatusBadRequest, nil)
	goto RETURN
RETURN:
	return
}
