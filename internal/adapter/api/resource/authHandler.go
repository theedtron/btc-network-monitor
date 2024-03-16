package resource

import (
	"btc-network-monitor/internal/adapter/api/requests"
	"btc-network-monitor/internal/adapter/api/response"
	"btc-network-monitor/internal/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *HTTPHandler) Register(c *gin.Context) {
	logger.Info("sign up user")

	request := requests.SignUpRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := s.auth.Register(request)
	if err != nil {
		logger.Error("Error saving sprint" + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.NewUserResponse(resp, err))
}

func (s *HTTPHandler) Login(c *gin.Context) {
	logger.Info("login user")

	request := requests.LoginRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := s.auth.Login(request)
	if err != nil {
		logger.Error("Error saving sprint" + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.NewUserResponse(resp, err))
}
