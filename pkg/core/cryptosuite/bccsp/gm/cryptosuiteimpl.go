/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package gm

import (
	"github.com/hyperledger/fabric-sdk-go/internal/github.com/hyperledger/fabric/bccsp"
	bccspGm "github.com/hyperledger/fabric-sdk-go/internal/github.com/hyperledger/fabric/bccsp/factory/gm"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/logging"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/core"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/cryptosuite/bccsp/wrapper"
	"github.com/pkg/errors"
)

var logger = logging.NewLogger("fabric_sdk_go")

//GetSuiteByConfig returns cryptosuite adaptor for bccsp loaded according to given config
func GetSuiteByConfig(config core.CryptoSuiteConfig) (core.CryptoSuite, error) {
	// TODO: delete this check?
	if config.SecurityProvider() != "gm" {
		return nil, errors.Errorf("Unsupported BCCSP Provider: %s", config.SecurityProvider())
	}

	opts := getOptsByConfig(config)
	bccsp, err := getBCCSPFromOpts(opts)
	if err != nil {
		return nil, err
	}
	return wrapper.NewCryptoSuite(bccsp), nil
}

//GetSuiteWithDefaultEphemeral returns cryptosuite adaptor for bccsp with default ephemeral options (intended to aid testing)
func GetSuiteWithDefaultEphemeral() (core.CryptoSuite, error) {
	opts := getEphemeralOpts()

	bccsp, err := getBCCSPFromOpts(opts)
	if err != nil {
		return nil, err
	}
	return wrapper.NewCryptoSuite(bccsp), nil
}

func getBCCSPFromOpts(config *bccspGm.GmOpts) (bccsp.BCCSP, error) {
	f := &bccspGm.GMFactory{}

	csp, err := f.Get(config)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not initialize BCCSP %s", f.Name())
	}
	return csp, nil
}

//GetOptsByConfig Returns Factory opts for given SDK config
func getOptsByConfig(c core.CryptoSuiteConfig) *bccspGm.GmOpts {
	opts := &bccspGm.GmOpts{
		HashFamily: c.SecurityAlgorithm(),
		SecLevel:   c.SecurityLevel(),
		FileKeystore: &bccspGm.FileKeystoreOpts{
			KeyStorePath: c.KeyStorePath(),
		},
		//Ephemeral: c.Ephemeral(),
	}
	logger.Debugf("Initialized GM cryptosuite, %v", c.SecurityAlgorithm())

	return opts
}

func getEphemeralOpts() *bccspGm.GmOpts {
	opts := &bccspGm.GmOpts{
		HashFamily: "SHA256",
		SecLevel:   256,
		Ephemeral:  true,
	}
	logger.Debug("Initialized ephemeral GM cryptosuite with default opts")

	return opts
}
