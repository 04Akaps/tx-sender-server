package node

import (
	"github.com/04Akaps/tx-sender-server/config"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

type NodeInfo struct {
	Client    *ethclient.Client
	URL       string
	Chain     string
	NetworkID *big.Int
}

type NodeClient struct {
	config *config.Config

	nodeInfos map[string]*NodeInfo
}

func NewNode(config *config.Config) (*NodeClient, error) {
	r := &NodeClient{
		config:    config,
		nodeInfos: make(map[string]*NodeInfo),
	}

	for chain, node := range config.Node {
		if client, err := ethclient.Dial(node.URL); err != nil {
			return nil, err
		} else {
			log.Println("Success To Connect Node", "URL", node.URL, "Chain", chain)
			r.nodeInfos[chain] = &NodeInfo{
				Client: client,
				URL:    node.URL,
				Chain:  chain,
			}
		}
	}

	return r, nil
}
