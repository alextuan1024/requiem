package v1

import (
	"github.com/alextuan1024/requiem/web"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

type Api struct {
}

var api *Api
var once sync.Once

func GetApi() *Api {
	once.Do(func() {
		api = &Api{}
	})
	return api
}

func (a *Api) handlerFoo(c *gin.Context) {
	c.String(http.StatusOK, "Bar")
}

func RegisterHandlers(e *gin.Engine, a *Api) {
	v1 := e.Group("/v1")
	v1.POST("/login", web.LoginHandler)

	v1a := v1.Use(web.AuthMiddleware)
	v1a.GET("/foo", a.handlerFoo)
}
