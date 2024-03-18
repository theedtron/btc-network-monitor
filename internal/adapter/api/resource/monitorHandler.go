package resource

import (
	"btc-network-monitor/internal/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *HTTPHandler) GetBlockchainInfo(c *gin.Context) {
	logger.Info("getBlockchainInfo")

	chainInfo, err := s.monitorService.GetBlockChainInfo()
	if err != nil {
		logger.Error("Error getting blockchain info: " + err.Error())
		c.JSON(500, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, chainInfo)

}

func (s *HTTPHandler) GetBlockByHash(c *gin.Context) {
	logger.Info("get block by hash")

	blockHashStr := c.Param("block_hash")
	block, err := s.monitorService.GetBlockByHash(blockHashStr)
	if err != nil {
		logger.Error("Error getting block by hash: " + err.Error())
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, block)
}
