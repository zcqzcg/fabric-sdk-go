// Copyright SecureKey Technologies Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

module github.com/hyperledger/fabric-sdk-go

require (
	github.com/Knetic/govaluate v3.0.1-0.20171022003610-9aa49832a739+incompatible
	github.com/VividCortex/gohistogram v1.0.0 // indirect
	github.com/cloudflare/cfssl v1.4.1
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/go-kit/kit v0.10.0
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/gogo/protobuf v1.2.1
	github.com/golang/mock v1.2.0
	github.com/golang/protobuf v1.4.0
	github.com/hyperledger/fabric-ca v1.4.6
	github.com/hyperledger/fabric-lib-go v1.0.0
	github.com/hyperledger/fabric-protos-go v0.0.0-20191121202242-f5500d5e3e85
	github.com/kr/pretty v0.1.0 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/miekg/pkcs11 v1.0.3
	github.com/mitchellh/mapstructure v1.2.2
	github.com/onsi/gomega v1.9.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.5.1
	github.com/spf13/cast v1.3.1
	github.com/spf13/viper v1.6.3
	github.com/stretchr/testify v1.5.1
	github.com/zcqzcg/gmsm v1.0.2
	github.com/zcqzcg/gmtls v1.0.0
	golang.org/x/crypto v0.0.0-20200406173513-056763e48d71
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e
	google.golang.org/grpc v1.28.1
	gopkg.in/yaml.v2 v2.2.8
)

replace (
	github.com/golang/protobuf v1.4.0 => github.com/golang/protobuf v1.3.2
	github.com/hyperledger/fabric-ca v1.4.6 => github.com/zcqzcg/fabric-ca v1.4.61
)

go 1.14
