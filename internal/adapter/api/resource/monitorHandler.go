package resource

import (
	"btc-network-monitor/internal/logger"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
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
	blockHash, err := chainhash.NewHashFromStr(blockHashStr)
	if err != nil {
		logger.Error("Invalid block hash: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	block, err := s.monitorService.GetBlockByHash(blockHash)
	if err != nil {
		logger.Error("Error getting block by hash: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, block)
}

func (s *HTTPHandler) GetTransaction(c *gin.Context) {
	logger.Info("get block by transaction id")

	txID := c.Param("tx_id")
	txHash, err := chainhash.NewHashFromStr(txID)
	if err != nil {
		logger.Error("Invalid transaction hash: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	tx, err := s.monitorService.GetTransactionByTransactionID(txHash)
	if err != nil {
		logger.Error("Error getting block by transaction id: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tx)
}

func (s *HTTPHandler) GetLatestTransactions(c *gin.Context) {
	logger.Info("get latest transactions")

	transactions, err := s.monitorService.GetLatestTransactions()
	if err != nil {
		logger.Error("Error getting latest transactions: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func (s *HTTPHandler) GetLatestBlocks(c *gin.Context) {
	logger.Info("get latest blocks")

	blocks, err := s.monitorService.GetLatestBlocks()
	if err != nil {
		logger.Error("Error getting latest blocks: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blocks)
}

func (s *HTTPHandler) GetAddressInfo(c *gin.Context) {
	logger.Info("get address info")

	address := c.Param("address")
	block, err := s.monitorService.GetAddressInfo(address)
	if err != nil {
		logger.Error("Error getting address info: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, block)
}
