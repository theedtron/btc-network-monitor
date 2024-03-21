package resource

import (
	"btc-network-monitor/internal/adapter/api/requests"
	"btc-network-monitor/internal/adapter/api/response"
	"btc-network-monitor/internal/core/domain"
	"btc-network-monitor/internal/logger"
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

func (s *HTTPHandler) CreateTxSubscribe(c *gin.Context) {
	logger.Info("Update Target Confirms")

	request := domain.TxSubscribe{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
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