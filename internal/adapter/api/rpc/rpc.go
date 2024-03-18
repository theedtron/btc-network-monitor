package rpc

import (
	"github.com/btcsuite/btcd/rpcclient"
	"os"
	"strconv"
)

func NewRPCConfig() {

	host := os.Getenv("RPC_HOST")
	user := os.Getenv("RPC_USER")
	pass := os.Getenv("RPC_PASSWORD")
	httpPostMode, _ := strconv.ParseBool(os.Getenv("HTTP_POST_MODE"))
	disableTLS, _ := strconv.ParseBool(os.Getenv("RPC_DISABLE_TLS"))

	config := &rpcclient.ConnConfig{
		Host:         host,
		User:         user,
		Pass:         pass,
		HTTPPostMode: httpPostMode,
		DisableTLS:   disableTLS,
	}

	var err error
	client, err := rpcclient.New(config, nil)
	if err != nil {
		panic(err)
	}

	Config = client

}

var Config *rpcclient.Client
