package resource

import (
	"btc-network-monitor/internal/adapter/api/requests"
	"btc-network-monitor/internal/adapter/api/response"
	"btc-network-monitor/internal/core/domain"
	"btc-network-monitor/internal/logger"
	"btc-network-monitor/internal/mailer"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *HTTPHandler) GetAllTxSubscribe(c *gin.Context) {
	logger.Info("GetAllTxSubscribers")

	params := FlatUrlQuery(c.Request.URL.Query())

	resp, err := s.txSubscribeService.GetAll(params)
	if err != nil {
		logger.Error("Error getting data" + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.NewTxSubscribeArrayResponse(resp, err))
}

func (s *HTTPHandler) GetFalseStatusSubscribers(c *gin.Context) {
	logger.Info("GetFalseStatusSubscribers")

	resp, err := s.txSubscribeService.GetFalseStatus()
	if err != nil {
		logger.Error("Error getting data" + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.NewTxSubscribeArrayResponse(resp, err))
}

func (s *HTTPHandler) CreateTxSubscribe(c *gin.Context) {
	logger.Info("Update Target Confirms")
	user, authExist := c.Get("user")
	if !authExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Authenticated user not detected",
		})
		return
	}
	userMap := user.(map[string]interface {})

	request := domain.TxSubscribe{UserID: userMap["id"].(string)}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//Check if exists
	_, err = s.txSubscribeService.FindByTxId(request.TxID)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Transaction already exists",
		})
		return
	}

	resp, err := s.txSubscribeService.Save(request)
	if err != nil {
		logger.Error("Error updating user" + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.NewTxSubscribeResponse(resp, err))
}

func (s *HTTPHandler) UpdateTxSubscribe(c *gin.Context) {
	logger.Info("Update Target Confirms")

	request := requests.UpdateTxSubscribeRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	params := c.Param("id")

	resp, err := s.txSubscribeService.Update(params, request)
	if err != nil {
		logger.Error("Error updating user" + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.NewTxSubscribeResponse(resp, err))
}

func (s *HTTPHandler) FindTxSubscribe(c *gin.Context) {
	logger.Info("Find Tx")

	id := c.Param("id")

	resp, err := s.txSubscribeService.Find(id)
	if err != nil {
		logger.Error("Error finding user" + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.NewTxSubscribeResponse(resp, err))
}

func (s *HTTPHandler) TestMail(c *gin.Context) {
	//send email
	config := mailer.NewNotificationConfig()
	senderEmailData := mailer.EmailData{
		FirstName:     "Theed",
		Subject:       "BTM Transaction Confirmation",
		MailTo:        "kareoedwin@gmail.com",
		Confirmations: 11,
		TxId:          "33adb5e349125eb07b74d2ef70a658e0bc3ca7bad7337600f91592d166559197",
	}

	config.SendEmail(&senderEmailData)

	c.JSON(http.StatusOK, "ok")
}
