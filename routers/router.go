package routers

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testserver/controllers"

	"testserver/models"

	"github.com/gin-gonic/gin"
	"github.com/littletwolee/commons"
)

var (
	Router *router
)

type router struct {
	Router *gin.Engine
	logger *commons.Log
}

func init() {
	if Router == nil {
		Router = &router{}
	}
	Router.logger = commons.GetLogger()
}

func (r *router) InitRouters(router *gin.Engine) {
	var (
		paths                       string
		err                         error
		routers                     *models.Router
		val, rou, controllersRouter reflect.Value
		mothedPaths                 []string
		souterMotheds               = []string{"Post", "Put", "Get", "Delete"}
	)
	paths = commons.Config.GetString("router.routers")
	if paths == "" {
		r.logger.LogPanic(fmt.Errorf("paths is empty"))
	}
	err = json.Unmarshal([]byte(paths), &routers)
	if err != nil {
		r.logger.LogPanic(err)
	}
	controllers.Router.Router = routers
	val = reflect.ValueOf(routers).Elem()
	rou = reflect.ValueOf(router)
	controllersRouter = reflect.ValueOf(controllers.Router)
	r.Router = router
	for _, v := range souterMotheds {
		mothedPaths = val.FieldByName(v).Interface().([]string)
		if mothedPaths != nil && len(mothedPaths) > 0 {
			for _, path := range mothedPaths {
				rou.MethodByName(strings.ToUpper(v)).Call([]reflect.Value{reflect.ValueOf(path), controllersRouter.MethodByName(v)})
			}
		}

	}
}
