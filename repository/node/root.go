package node

import (
	"context"
	"github.com/04Akaps/tx-sender-server/config"
	. "github.com/04Akaps/tx-sender-server/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

type NodeInfo struct {
	client          *ethclient.Client
	url             string
	chain           string
	networkID       *big.Int
	chainID         *big.Int
	signer          types.Signer
	defaultGasLimit uint64
}

type NodeClient struct {
	config *config.Config

	nodeInfos map[string]*NodeInfo
}

type NodeImpl interface {
	GetCodeAt(chain, address string) ([]byte, error)
	GetBalance(chain, address string) (*big.Int, error)
	GetLatestBlockNumber(chain string) (uint64, error)
	GetBlockByNumber(chain string, number *big.Int) (*types.Block, error)
	GetTxReceipt(chain string, tx string) (*types.Receipt, error)
	GetSigner(chain string) (types.Signer, error)
	GetNonce(chain, from string) (uint64, error)
	GetFeePerGas(chain string) (*big.Int, error)
	GetBaseFee(chain string) (*big.Int, error)
	GetChainID(chain string) (*big.Int, error)
	GetDefaultGasLimit(chain string) (uint64, error)
	SendTransaction(chain string, tx *types.Transaction) error
}

func NewNode(config *config.Config) (NodeImpl, error) {
	r := &NodeClient{
		config:    config,
		nodeInfos: make(map[string]*NodeInfo),
	}

	for chain, node := range config.Node {
		if client, err := ethclient.Dial(node.URL); err != nil {
			return nil, err
		} else if chainID, err := client.ChainID(context.Background()); err != nil {
			return nil, err
		} else if networkID, err := client.NetworkID(context.Background()); err != nil {
			return nil, err
		} else {
			log.Println("Success To Connect Node", "URL", node.URL, "Chain", chain)
			r.nodeInfos[chain] = &NodeInfo{
				client:    client,
				url:       node.URL,
				chain:     chain,
				networkID: networkID,
				chainID:   chainID,
				signer:    types.NewLondonSigner(chainID),
			}
		}
	}

	return r, nil
}

func (n *NodeClient) getClient(chain string) (*ethclient.Client, error) {
	if nodeInfo, err := n.getNodeInfo(chain); err != nil {
		return nil, err
	} else {
		return nodeInfo.client, nil
	}
}

func (n *NodeClient) getNodeInfo(chain string) (*NodeInfo, error) {
	if node, ok := n.nodeInfos[chain]; !ok {
		return nil, ToNodeErrType(InvalidNodeInfo)
	} else {
		return node, nil
	}
}

func (n *NodeClient) getPendingNonce(chain, from string) (uint64, error) {
	if client, err := n.getClient(chain); err != nil {
		return 0, err
	} else if res, err := client.PendingNonceAt(context.Background(), toAddress(from)); err != nil {
		return 0, err
	} else {
		return res, nil
	}
}

func (n *NodeClient) SendTransaction(chain string, tx *types.Transaction) error {
	if client, err := n.getClient(chain); err != nil {
		return err
	} else if err = client.SendTransaction(context.Background(), tx); err != nil {
		return err
	} else {
		return nil
	}
}

func (n *NodeClient) GetNonce(chain, from string) (uint64, error) {
	if nonce, err := n.getPendingNonce(chain, from); err != nil {
		return 0, err
	} else {
		return nonce, nil
	}
}

func (n *NodeClient) GetSigner(chain string) (types.Signer, error) {
	if node, err := n.getNodeInfo(chain); err != nil {
		return nil, err
	} else {
		return node.signer, nil
	}
}

func (n *NodeClient) GetBalance(chain, address string) (*big.Int, error) {
	if client, err := n.getClient(chain); err != nil {
		return nil, err
	} else if balance, err := client.BalanceAt(context.Background(), toAddress(address), nil); err != nil {
		return nil, err
	} else {
		return balance, nil
	}
}

func (n *NodeClient) GetLatestBlockNumber(chain string) (uint64, error) {
	if client, err := n.getClient(chain); err != nil {
		return 0, err
	} else if num, err := client.BlockNumber(context.Background()); err != nil {
		return 0, err
	} else {
		return num, nil
	}
}

func (n *NodeClient) GetBlockByNumber(chain string, number *big.Int) (*types.Block, error) {
	if client, err := n.getClient(chain); err != nil {
		return nil, err
	} else if res, err := client.BlockByNumber(context.Background(), number); err != nil {
		return nil, err
	} else {
		return res, nil
	}
}

func (n *NodeClient) GetTxReceipt(chain string, tx string) (*types.Receipt, error) {
	if client, err := n.getClient(chain); err != nil {
		return nil, err
	} else if res, err := client.TransactionReceipt(context.Background(), toHash(tx)); err != nil {
		return nil, err
	} else {
		return res, nil
	}
}

func (n *NodeClient) GetCodeAt(chain, address string) ([]byte, error) {
	if client, err := n.getClient(chain); err != nil {
		return nil, err
	} else if bytes, err := client.CodeAt(context.Background(), toAddress(address), nil); err != nil {
		return nil, err
	} else {
		return bytes, nil
	}
}

func (n *NodeClient) GetFeePerGas(chain string) (*big.Int, error) {
	if client, err := n.getClient(chain); err != nil {
		return nil, err
	} else if res, err := client.SuggestGasTipCap(context.Background()); err != nil {
		return nil, err
	} else {
		return res, nil
	}
}

func (n *NodeClient) GetChainID(chain string) (*big.Int, error) {
	if client, err := n.getNodeInfo(chain); err != nil {
		return nil, err
	} else {
		return client.chainID, nil
	}
}

func (n *NodeClient) GetBaseFee(chain string) (*big.Int, error) {
	if number, err := n.GetLatestBlockNumber(chain); err != nil {
		return nil, err
	} else if block, err := n.GetBlockByNumber(chain, big.NewInt(int64(number))); err != nil {
		return nil, err
	} else {
		return block.BaseFee(), nil
	}
}

func (n *NodeClient) GetDefaultGasLimit(chain string) (uint64, error) {
	if client, err := n.getNodeInfo(chain); err != nil {
		return 0, err
	} else {
		return client.defaultGasLimit, nil
	}
}

func toAddress(str string) common.Address {
	return common.HexToAddress(str)
}

func toHash(str string) common.Hash {
	return common.HexToHash(str)
}
