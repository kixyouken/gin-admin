package delivery

import (
	"gin-admin/common/output"
	"gin-admin/model"
	"gin-admin/param"
	"gin-admin/search"
	"gin-admin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	modelDeliveryList []model.MailDeliveryModel
	paramDeliveryList []param.MailDeliveryInfo
	err               error
)

func Index(c *gin.Context) {
	menuTree, _ := c.Get("menuTree")
	c.HTML(http.StatusOK, "delivery/index", gin.H{
		"title": "メール配信管理",
		"menu":  menuTree,
		"uri":   "mailer/delivery",
	})
}

func Data(c *gin.Context) {
	paramMailDeliverySearch := param.MailDeliverySearch{}
	c.ShouldBind(&paramMailDeliverySearch)
	searchDB := search.DeliverySearch.SearchDelivery(paramMailDeliverySearch)

	count := service.MailDeliveryService.GetCountDelivery(c, searchDB)
	if count > 0 {
		modelDeliveryList, err = service.MailDeliveryService.GetAllDelivery(c, searchDB, "")
		if err != nil {
			return
		}
		paramDeliveryList = service.MailDeliveryService.DeliveryList(modelDeliveryList)
	} else {
		paramDeliveryList = nil
	}

	output.Pageful(c, paramDeliveryList, count)
}
