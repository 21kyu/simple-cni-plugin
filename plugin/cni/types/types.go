package types

import (
	"encoding/json"
	"fmt"

	"github.com/containernetworking/cni/pkg/types"
	"github.com/containernetworking/cni/pkg/version"
)

const defaultBrName = "cni0"

type NetConf struct {
	types.NetConf

	// Add plugin-specifc flags here
	BrName string `json:"bridge"`
}

// LoadNetConf unmarshals the network config from stdin and returns
func LoadNetConf(stdin []byte) (*NetConf, error) {
	conf := &NetConf{
		BrName: defaultBrName,
	}

	if err := json.Unmarshal(stdin, conf); err != nil {
		return nil, fmt.Errorf("failed to parse network configuration: %v", err)
	}

	if err := version.ParsePrevResult(&conf.NetConf); err != nil {
		return nil, fmt.Errorf("could not parse prevResult: %v", err)
	}

	// Do any validation here
	//

	return conf, nil
}
