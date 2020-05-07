package cli

import (
	"errors"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/context"
)

type FabricUser struct {
	Username string
	Secret string
	Affiliation string
	Type string
}

func (c *Client)RegisterUser(username string, password string, affiliation string, attributes []msp.Attribute) (string, error) {

	mspClient := c.MspClient
	//rsp,err := mspClient.GetCAInfo()
	//if err != nil {
	//	log.Panic("xxxxxxxxxxxxxxxnot able to getCAInfo,",err)
	//}
	//b,_:=json.Marshal(rsp)
	//log.Info("xxxxxxxxxxCA INFO",string(b))

	instance, s := getRegistrarEnrollmentCredentialsWithCAInstance(c.SDK.Context(), "")

	//登陆registrar
	err := mspClient.Enroll(instance, msp.WithSecret(s))
	if err != nil {
		return "", err
	}

	register := msp.RegistrationRequest{
		Name:        username,
		//使用Type可能导致ca注册的用户无法建立通道
		//Type:        "user",
		Affiliation: affiliation,
		Attributes:  attributes,
		Secret:      password,
	}
	//注册用户，secret为登陆密钥
	secret, err := mspClient.Register(&register)
	if err != nil {
		return "", err
	}
	return secret, nil
}

func (c *Client) EnrollUser(username,s string) error {
	if username == "" || s == "" {
		return errors.New("we need a username/secret to enroll")
	}
	mspClient := c.MspClient
	return mspClient.Enroll(username,msp.WithSecret(s))
}

func getRegistrarEnrollmentCredentialsWithCAInstance( ctxProvider context.ClientProvider, caID string) (string, string) {
	ctx, err := ctxProvider()
	if err != nil {
		panic(err)
	}
	myOrg := ctx.IdentityConfig().Client().Organization
	//mgr,_ :=ctx.IdentityConfig().Client().
	if caID == "" {
		if len(ctx.EndpointConfig().NetworkConfig().Organizations[myOrg].CertificateAuthorities) != 0 {
			caID = ctx.EndpointConfig().NetworkConfig().Organizations[myOrg].CertificateAuthorities[0]
		} else {
			panic("we don't have a correct caID, please check ccp file.")
		}
	}
	caConfig, ok := ctx.IdentityConfig().CAConfig(caID)
	if !ok {
		panic(err)
	}
	return caConfig.Registrar.EnrollID, caConfig.Registrar.EnrollSecret
}
