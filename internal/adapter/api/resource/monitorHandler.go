package resource

import (
	"btc-network-monitor/internal/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func (s *HTTPHandler) GetBlockByHeight(c *gin.Context) {
	logger.Info("get block by height")

	blockHeightStr := c.Param("param")
	blockHeight, err := strconv.ParseInt(blockHeightStr, 10, 32)
	if err != nil {
		logger.Error("Invalid block height: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid block height"})
		return
	}
	block, err := s.monitorService.GetBlockByHeight(blockHeight)
	if err != nil {
		logger.Error("Error getting block by height: " + err.Error())
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, block)
}

func (s *HTTPHandler) GetTransaction(c *gin.Context) {
	logger.Info("get block by transaction id")

	txID := c.Param("tx_id")
	tx, err := s.monitorService.GetTransactionByTransactionID(txID)
	if err != nil {
		logger.Error("Error getting block by transaction id: " + err.Error())
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tx)
}
