package resource

import (
	"btc-network-monitor/internal/adapter/api/requests"
	"btc-network-monitor/internal/adapter/api/response"
	"btc-network-monitor/internal/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *HTTPHandler) GetAllUsers(c *gin.Context) {
	logger.Info("GetAllUsers")

	params := FlatUrlQuery(c.Request.URL.Query())

	resp, err := s.userService.GetAll(params)
	if err != nil {
		logger.Error("Error getting data" + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.NewUserArrayResponse(resp, err))
}

func (s *HTTPHandler) UpdateUser(c *gin.Context) {
	logger.Info("Update User")

	request := requests.UpdateUserRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	params := c.Param("id")

	resp, err := s.userService.Update(params, request)
	if err != nil {
		logger.Error("Error updating user" + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.NewUserResponse(resp, err))
}

func (s *HTTPHandler) FindUser(c *gin.Context) {
	logger.Info("Find User")

	id := c.Param("id")

	resp, err := s.userService.Find(id)
	if err != nil {
		logger.Error("Error finding user" + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.NewUserResponse(resp, err))
}
