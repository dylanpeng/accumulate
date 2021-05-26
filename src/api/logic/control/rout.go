package control

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var Rout = routCtrl{}

type routCtrl struct{}

func (c *routCtrl) GetParametersInPath(ctx *gin.Context) {
	name := ctx.Param("name")
	action := ctx.Param("action")
	message := name + " is " + action
	ctx.String(http.StatusOK, message)
}
