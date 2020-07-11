package v1

import (
	"github.com/gin-gonic/gin"
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

func RegisterHandlers(e *gin.Engine, a *Api) {
}
