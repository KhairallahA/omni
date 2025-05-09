package geth

import (
	"net"
	"reflect"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p/enode"

	"github.com/naoina/toml"
)

// ServerVersion defines the geth version deployed to all networks.
const ServerVersion = "v1.14.13"

// ClientVersion defines the geth version used mostly for ethclient.
const ClientVersion = "v1.15.6"

// SupportedVersions are the supported older geth server versions.
// These are tested in backwards.toml.
var SupportedVersions = []string{
	"v1.14.12",
	"v1.14.11",
	"v1.14.8",
	"v1.14.7",
}

// Config is the configurable options for the standard omni geth config.
type Config struct {
	// Moniker is the p2p node name.
	Moniker string
	// ChainID is the chain ID of the network.
	ChainID uint64
	// IsArchive defines whether the node should run in archive mode.
	IsArchive bool
	// BootNodes are the enode URLs of the P2P bootstrap nodes.
	BootNodes []*enode.Node
	// TrustedNodes are the enode URLs of the P2P trusted nodes.
	TrustedNodes []*enode.Node
	// SnapshotCacheMB overrides the default snapshot cache size in MB if not zero.
	SnapshotCacheMB int
	// AdvertisedIP
	AdvertisedIP net.IP
}

// defaultGethConfig returns the default geth config.
func defaultGethConfig() FullConfig {
	return FullConfig{
		Eth:     ethconfig.Defaults,
		Node:    node.DefaultConfig,
		Metrics: metrics.DefaultConfig, // Enable prometheus metrics via command line flags --metrics --pprof --pprof.addr=0.0.0.0
	}
}

// FullConfig is the go struct representation of the geth.toml config file.
// Copied from https://github.com/ethereum/go-ethereum/blob/master/cmd/geth/config.go#L95
type FullConfig struct {
	Eth     ethconfig.Config
	Node    node.Config
	Metrics metrics.Config
}

// tomlSettings is the toml settings used to parse/format the geth.toml config file.
// Copied from https://github.com/ethereum/go-ethereum/blob/master/cmd/geth/config.go#L70.
var tomlSettings = toml.Config{
	NormFieldName: func(_ reflect.Type, key string) string {
		return key
	},
	FieldToKey: func(_ reflect.Type, field string) string {
		return field
	},
	MissingField: func(rt reflect.Type, field string) error {
		return errors.New("field not defined", "field", field, "type", rt.String())
	},
}
