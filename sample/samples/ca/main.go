package main

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"log"

	"github.com/hyperledger/fabric-sdk-go/sample/cli"
)

const (
	org1CfgPath = "/home/david/go/src/github.com/hyperledger/fabric-sdk-go/sample/config/org1sdk-config.yaml"
	org2CfgPath = "/home/david/go/src/github.com/hyperledger/fabric-sdk-go/sample/config/org2sdk-config.yaml"
)

func main() {
	log.Println("Starting test ca.")

	org1Client := cli.New(org1CfgPath, "Org1", "Admin", "User1")
	//org2Client := cli.New(org2CfgPath, "Org2", "Admin", "User1")
	//
	defer org1Client.Close()
	//defer org2Client.Close()

	// Enroll admin; Reg New User; Enroll user
	username := Phase1(org1Client, nil)
	// Using user to query
	Phase2(org1Client, nil,username)
}

// Enroll admin; Reg New User; Enroll user
func Phase1(cli1, cli2 *cli.Client) (username string){
	log.Println("=================== Phase 1 begin ===================")
	defer log.Println("=================== Phase 1 end ===================")
	userName := "goUser6"
	userSecret := "123456"
	//Enroll admin; Reg New User;
	secret,err := cli1.RegisterUser(userName,userSecret,"",nil)
	if err != nil {
		log.Panicf("RegisterUser error: %v", err)
	}

	//Enroll user
	err = cli1.EnrollUser(userName,secret)
	if err != nil {
		log.Panicf("EnrollUser error: %v", err)
	}

	return userName
}

func Phase2(cli1, cli2 *cli.Client,username string) {
	log.Println("=================== Phase 2 begin ===================")
	defer log.Println("=================== Phase 2 end ===================")

	cli1.SDK.Context(fabsdk.WithUser(username))

	log.Println("Change SDK context User to:",username)
	if _, err := cli1.InvokeCC([]string{"peer0.org1.example.com",
		"peer0.org2.example.com"}); err != nil {
		log.Panicf("Invoke chaincode error: %v", err)
	}
	log.Println("Invoke chaincode success")

	if err := cli1.QueryCC("peer0.org2.example.com", "a"); err != nil {
		log.Panicf("Query chaincode error: %v", err)
	}
	log.Println("Query chaincode success on peer0.org2")
}
