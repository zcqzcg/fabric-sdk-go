package cli

import (
	//"log"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/logging"
	"os"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

var log = logging.NewLogger("sample/cli/client")

type Client struct {
	// Fabric network information
	ConfigPath string
	OrgName    string
	OrgAdmin   string
	OrgUser    string

	// sdk clients
	SDK       *fabsdk.FabricSDK
	rc        *resmgmt.Client
	cc        *channel.Client
	MspClient *msp.Client

	// Same for each peer
	ChannelID string
	CCID      string // chaincode ID, eq name
	CCPath    string // chaincode source path, 是GOPATH下的某个目录
	CCGoPath  string // GOPATH used for chaincode
}

func New(cfg, org, admin, user string) *Client {
	c := &Client{
		ConfigPath: cfg,
		OrgName:    org,
		OrgAdmin:   admin,
		OrgUser:    user,

		CCID:      "mycc",
		CCPath:    "github.com/hyperledger/fabric-samples-gm/chaincode/chaincode_example02/go/", // 相对路径是从GOPAHT/src开始的
		CCGoPath:  os.Getenv("GOPATH"),
		ChannelID: "mychannel",
	}

	// create sdk
	log.Println(c.ConfigPath)
	sdk, err := fabsdk.New(config.FromFile(c.ConfigPath))
	if err != nil {
		log.Panicf("failed to create fabric sdk: %s", err)
	}
	c.SDK = sdk
	log.Println("Initialized fabric sdk")

	c.rc, c.cc,c.MspClient = NewSdkClient(sdk, c.ChannelID, c.OrgName, c.OrgAdmin, c.OrgUser)

	return c
}

// NewSdkClient create resource client and channel client
func NewSdkClient(sdk *fabsdk.FabricSDK, channelID, orgName, orgAdmin, OrgUser string) (rc *resmgmt.Client, cc *channel.Client, mspClient *msp.Client) {
	var err error

	// create rc
	rcp := sdk.Context(fabsdk.WithUser(orgAdmin), fabsdk.WithOrg(orgName))
	rc, err = resmgmt.New(rcp)
	if err != nil {
		log.Panicf("failed to create resource client: %s", err)
	}
	log.Println("Initialized resource client")

	// create cc
	ccp := sdk.ChannelContext(channelID, fabsdk.WithUser(OrgUser))
	cc, err = channel.New(ccp)
	if err != nil {
		log.Panicf("failed to create channel client: %s", err)
	}
	log.Println("Initialized channel client")

	// create mspClient
	mspClient, err = msp.New(rcp)
	if err != nil {
		log.Panicf("failed to create channel client: %s", err)
	}
	log.Println("Initialized msp client")

	return rc, cc, mspClient
}

// RegisterChaincodeEvent more easy than event client to registering chaincode event.
func (c *Client) RegisterChaincodeEvent(ccid, eventName string) (fab.Registration, <-chan *fab.CCEvent, error) {
	return c.cc.RegisterChaincodeEvent(ccid, eventName)
}
