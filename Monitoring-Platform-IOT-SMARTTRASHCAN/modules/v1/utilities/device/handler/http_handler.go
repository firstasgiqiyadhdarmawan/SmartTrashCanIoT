package product

import (
	"ClearningPatternGO/modules/v1/utilities/device/models"
	"ClearningPatternGO/modules/v1/utilities/device/repository"
	"ClearningPatternGO/modules/v1/utilities/device/service"
	api "ClearningPatternGO/pkg/api_response"
	"ClearningPatternGO/pkg/helpers"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DeviceHandler interface {
	ListProduct(c *gin.Context)
}

type deviceHandler struct {
	productService service.Service
}

func NewDeviceHandler(productService service.Service) *deviceHandler {
	return &deviceHandler{productService}
}

func Handler(db *gorm.DB) *deviceHandler {
	Repository := repository.NewRepository(db)
	Service := service.NewService(Repository)
	Handler := NewDeviceHandler(Service)
	return Handler
}

func (h *deviceHandler) ListDevice(c *gin.Context) {
	listProduct, err := h.productService.ListDevice()

	if err != nil {
		response := api.APIRespon("Failed to get data all device", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := api.APIRespon("Success to get data all device", http.StatusOK, "success", listProduct)
	c.JSON(http.StatusOK, response)
}

func (n *deviceHandler) ReceivedDataAntares(c *gin.Context) {
	access_key := "2c4b04aa13dbc9bc:cb8f6d670f09d6fc"
	getLatestContent, err := n.productService.GetLatestContent(access_key)
	if err != nil {
		fmt.Println(err)
		return
	}
	Antares_Device_Id := strings.Replace(getLatestContent.First.Pi, "/antares-cse/cnt-", "", -1)
	fmt.Println("ini Antares Device ID nya \n", Antares_Device_Id)
	_, err = n.productService.GetDatafromContent(getLatestContent.First.Con, Antares_Device_Id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Sukses masuk data ke db")
	c.JSON(http.StatusOK, getLatestContent.First.Con)
}

func (n *deviceHandler) SubscribeWebhook(c *gin.Context) {
	var webhookData models.ObjectAntares1
	if err := c.ShouldBindJSON(&webhookData); err != nil {
		response := helpers.APIRespon("Error, inputan tidak sesuai", 220, "error", nil)
		c.JSON(220, response)
		return
	}
	Antares_Device_Id := strings.Replace(webhookData.First.M2m_nev.M2m_rep.M2m_cin.Pi, "/antares-cse/cnt-", "", -1)
	_, err := n.productService.GetDatafromWebhook(webhookData.First.M2m_nev.M2m_rep.M2m_cin.Con, Antares_Device_Id)
	if err != nil {
		fmt.Println(err)
		return
	}
}
