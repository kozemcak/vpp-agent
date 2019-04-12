// Copyright (c) 2017 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"os"
	"strings"

	"github.com/gogo/protobuf/proto"

	"fmt"

	"github.com/ligato/cn-infra/db/keyval"
	"github.com/ligato/cn-infra/db/keyval/etcd"
	"github.com/ligato/cn-infra/db/keyval/kvproto"
	"github.com/ligato/cn-infra/logging"
	"github.com/ligato/cn-infra/logging/logrus"
	"github.com/ligato/cn-infra/servicelabel"
)

const (
	ACLPath          = "config/vpp/acls/v2/acl/"
	InterfacePath    = "config/vpp/v2/interfaces/"
	BridgeDomainPath = "config/vpp/l2/v2/bridge-domain/"
	FibTablePath     = "config/vpp/l2/v2/fib/"
	XConnectPath     = "config/vpp/l2/v2/xconnect/"
	ARPPath          = "config/vpp/v2/arp/"
	RoutePath        = "config/vpp/v2/route/"
	ProxyARPPath     = "config/vpp/v2/proxyarp-global"
	IPScanneightPath = "config/vpp/v2/ipscanneigh-global"
	NATPath          = "config/vpp/nat/v2/nat44-global"
	DNATPath         = "config/vpp/nat/v2/dnat44/"
	IPSecPolicyPath  = "config/vpp/ipsec/v2/spd/"
	IPSecAssociate   = "config/vpp/ipsec/v2/sa/"
)

const (
	lInterfacePath = "config/linux/interfaces/v2/interface/"
	lARPPath       = "config/linux/l4/v2/arp/"
	lRoutePath     = "config/linux/l3/v2/route/"
)

// Common exit flags
const (
	ExitSuccess = iota
	ExitError
	ExitBadConnection
	ExitInvalidInput
	ExitBadFeature
	ExitInterrupted
	ExitIO
	ExitBadArgs = 128
)

// ParseKey parses the etcd Key for the microservice label and the
// data type encoded in the Key. The function returns the microservice
// label, the data type and the list of parameters, that contains path
// segments that follow the data path segment in the Key URL. The
// parameter list is empty if data path is the Last segment in the
// Key.
//

func ParseKey(key string) (label string, dataType string, name string) {
	ps := strings.Split(strings.TrimPrefix(key, servicelabel.GetAllAgentsPrefix()), "/")
	var plugin, statusConfig, version string
	var params []string
	if len(ps) > 0 {
		label = ps[0]
	}
	if len(ps) > 1 {
		plugin = ps[1]
		dataType = plugin
	}
	if len(ps) > 2 {
		statusConfig = ps[2]
		dataType += "/" + statusConfig

		if "vpp" == ps[2] {
			if len(ps) > 3 {
				version = ps[3]
				dataType += "/" + version
			}

			// Recognize key type
			if "v2" == ps[3] {
				if len(ps) > 4 {
					tp := ps[4]
					dataType += "/" + tp
				}

				if len(ps) > 5 {
					dataType += "/"
					params = ps[5:]
				} else {
					params = []string{}
				}
			} else {
				if len(ps) > 4 {
					version := ps[4]
					dataType += "/" + version
				}

				if len(ps) > 5 {
					tp := ps[5]
					dataType += "/" + tp
				}

				if len(ps) > 6 {
					dataType += "/"
					params = ps[6:]
				} else {
					params = []string{}
				}
			}
		} else if "linux" == ps[2] {
			if len(ps) > 3 {
				version = ps[3]
				dataType += "/" + version
			}

			if len(ps) > 4 {
				version := ps[4]
				dataType += "/" + version
			}

			if len(ps) > 5 {
				tp := ps[5]
				dataType += "/" + tp
			}

			if len(ps) > 6 {
				dataType += "/"
				params = ps[6:]
			} else {
				params = []string{}
			}
		} else {
			params = []string{}
		}
	} else {
		params = []string{}
	}

	return label, dataType, rebuildName(params)
}

// Reconstruct item name in case it contains slashes.
func rebuildName(params []string) string {
	var itemName string
	if len(params) > 1 {
		for _, param := range params {
			itemName = itemName + "/" + param
		}
		// Remove the first slash.
		return itemName[1:]
	} else if len(params) == 1 {
		itemName = params[0]
		return itemName
	}
	return itemName
}

// GetDbForAllAgents opens a connection to etcd, specified in the command line
// or the "ETCD_ENDPOINTS" environment variable.
func GetDbForAllAgents(endpoints []string) (keyval.ProtoBroker, error) {
	if len(endpoints) > 0 {
		ep := strings.Join(endpoints, ",")
		os.Setenv("ETCD_ENDPOINTS", ep)
	}

	cfg := &etcd.Config{}
	etcdConfig, err := etcd.ConfigToClient(cfg)

	// Log warnings and errors only.
	log := logrus.DefaultLogger()
	log.SetLevel(logging.WarnLevel)
	etcdBroker, err := etcd.NewEtcdConnectionWithBytes(*etcdConfig, log)
	if err != nil {
		return nil, err
	}

	return kvproto.NewProtoWrapperWithSerializer(etcdBroker, &keyval.SerializerJSON{}), nil

}

func GetModuleName(module proto.Message) string {
	str := proto.MessageName(module)

	tmp := strings.Split(str, ".")

	outstr := tmp[len(tmp)-1]

	if "linux" == tmp[0] {
		outstr = "Linux" + outstr
	}

	return outstr
}

// ExitWithError is used by all commands to print out an error
// and exit.
func ExitWithError(code int, err error) {
	fmt.Fprintln(os.Stderr, "Error: ", err)
	os.Exit(code)
}