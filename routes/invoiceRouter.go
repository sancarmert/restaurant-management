package routes

import (
	controller "github.com/sancarmert/restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/invoices", controller.GetInvoices())
	incomingRoutes.GET("/invoices/:invoice_id", controller.GetInvoice())
	incomingRoutes.POST("/invoices", controller.CreateInvoices())
	incomingRoutes.PATCH("/invoices/:invoice_id", controller.UpdateInvoices())

}
