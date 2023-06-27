package routes

import (
	"ClearningPatternGO/app/config"
	deviceHandlerV1 "ClearningPatternGO/modules/v1/utilities/device/handler"
	deviceviewV1 "ClearningPatternGO/modules/v1/utilities/device/view"
	basic "ClearningPatternGO/pkg/basic_auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ParseTmpl(router *gin.Engine) *gin.Engine { //Load HTML Template
	router.Static("/assets", "./public/assets")
	router.Static("/images", "./public/images")
	router.Static("/css", "./public/assets/css")
	router.Static("/js", "./public/assets/js")
	router.Static("/fonts", "./public/assets/fonts")
	return router
}

func Init(db *gorm.DB, conf config.Conf, router *gin.Engine) *gin.Engine {
	deviceHandlerV1 := deviceHandlerV1.Handler(db)
	deviceViewV1 := deviceviewV1.View(db)

	// Routing Website Service
	device := router.Group("/", basic.Auth(conf))
	device.GET("/", deviceViewV1.Index)

	//Routing API Service
	api := router.Group("/api/v1")
	api.GET("/device", deviceHandlerV1.ListDevice)
	api.GET("/antares-latest-data", deviceHandlerV1.ReceivedDataAntares)
	api.POST("/webhook", deviceHandlerV1.SubscribeWebhook)
	router = ParseTmpl(router)
	return router
}
