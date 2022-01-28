package builder

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/simapp/params"

	"github.com/pylons-tech/juno/node"
	nodeconfig "github.com/pylons-tech/juno/node/config"
	"github.com/pylons-tech/juno/node/local"
	"github.com/pylons-tech/juno/node/remote"
)

func BuildNode(cfg nodeconfig.Config, encodingConfig *params.EncodingConfig) (node.Node, error) {
	switch cfg.Type {
	case nodeconfig.TypeRemote:
		return remote.NewNode(cfg.Details.(*remote.Details), encodingConfig.Marshaler)
	case nodeconfig.TypeLocal:
		return local.NewNode(cfg.Details.(*local.Details), encodingConfig.TxConfig, encodingConfig.Marshaler)

	default:
		return nil, fmt.Errorf("invalid node type: %s", cfg.Type)
	}
}
