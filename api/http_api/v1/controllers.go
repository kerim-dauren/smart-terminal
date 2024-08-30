package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/kerim-dauren/smart-terminal/api/http_api/v1/controller"
	"github.com/kerim-dauren/smart-terminal/internal/manager"
)

func Register(routerGroup *gin.RouterGroup, managers *manager.Managers) {

	paymentController := controller.NewPaymentController(managers.KaspiPaymentManager)
	routerGroup.GET("/kaspi/payment", paymentController.ProcessKaspiRequest)
}
