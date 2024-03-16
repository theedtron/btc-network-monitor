package rpc

import (
	"github.com/btcsuite/btcd/rpcclient"
	"os"
	"strconv"
)

func NewRPCConfig() *rpcclient.ConnConfig {
	host := os.Getenv("RPC_HOST")
	user := os.Getenv("RPC_USER")
	pass := os.Getenv("PASSWORD")
	httpPostMode, _ := strconv.ParseBool(os.Getenv("HTTP_POST_MODE"))
	disableTLS, _ := strconv.ParseBool(os.Getenv("RPC_DISABLE_TLS"))

	return &rpcclient.ConnConfig{
		Host:         host,
		User:         user,
		Pass:         pass,
		HTTPPostMode: httpPostMode,
		DisableTLS:   disableTLS,
	}
}

var rpcConfig = NewRPCConfig()
