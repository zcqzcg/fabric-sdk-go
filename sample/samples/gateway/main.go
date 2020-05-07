package main

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/common/logging"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

const (
	org1CfgPath = "/home/david/go/src/github.com/hyperledger/fabric-sdk-go/sample/config/org1sdk-config.yaml"
	org2CfgPath = "/home/david/go/src/github.com/hyperledger/fabric-sdk-go/sample/config/org2sdk-config.yaml"
)

var logger = logging.NewLogger("gatewaySample")

/*
 * this is a sample file to test the using of
 * gateway package to get our hands on fabric
*/

func main() {
	username := "goUser6"

	wallet,err := gateway.NewFileSystemWallet("/tmp/examplestore")
	if err != nil {
		logger.Panic("new wallet failed::",err)
	}
	if wallet.Exists(username) {
		logger.Info(username," exists in wallet.")
	} else {
		logger.Info(username," not enrolled yet.")

	}


	myGateway,err := gateway.Connect(gateway.WithConfig(config.FromFile(org1CfgPath)),gateway.WithDiscovery(false),gateway.WithUser(username))
	if err != nil {
		logger.Panic("gateway connect error.",err)
	}
	defer myGateway.Close()
	logger.Info("Gateway Connected.")



	network,err := myGateway.GetNetwork("mychannel")
	if err != nil {
		logger.Panic("myGateway.GetNetwork error.",err)
	}
	logger.Info("Network Got::",network.Name())



	cc := network.GetContract("mycc")
	logger.Info("Got cc::",cc.Name())

}