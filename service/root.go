package service

import (
	"fmt"
	"github.com/04Akaps/tx-sender-server/repository/node"
	"github.com/04Akaps/tx-sender-server/repository/redis"
	"github.com/04Akaps/tx-sender-server/types"
	etypes "github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
)

type Service struct {
	node  node.NodeImpl
	redis redis.RedisImpl
}

func NewService(node node.NodeImpl, redis redis.RedisImpl) *Service {
	return &Service{node, redis}
}

func (s *Service) UnSignTx(req types.UnSignReq) {
	if code, err := s.getCode(req.Chain, req.Address); err != nil {
		log.Println(err)
	} else {
		fmt.Println(code)

		if len(code) > 0 {
			// Contract
		}

	}
}

func (s *Service) GetLatestBlockNumber(chain string) (uint64, error) {
	if res, err := s.node.GetLatestBlockNumber(chain); err != nil {
		log.Println("Failed To Get latestBlock", "chain", chain)
		return 0, err
	} else {
		return res, nil
	}
}

func (s *Service) GetBalance(chain, address string) (int64, error) {
	if res, err := s.node.GetBalance(chain, address); err != nil {
		log.Println("Failed To Get Balance", "chain", chain, "address", address)
		return 0, err
	} else {
		return res.Int64(), nil
	}
}

func (s *Service) GetBlockByNumber(chain string, number int64) (*etypes.Block, error) {
	if res, err := s.node.GetBlockByNumber(chain, big.NewInt(number)); err != nil {
		log.Println("Failed To Get Block", "chain", chain, "number", number)
		return nil, err
	} else {
		return res, nil
	}
}

func (s *Service) GetTxReceipt(chain, tx string) (*etypes.Receipt, error) {
	if res, err := s.node.GetTxReceipt(chain, tx); err != nil {
		log.Println("Failed To Get Tx Receipt", "chain", chain, "tx", tx)
		return nil, err
	} else {
		return res, nil
	}
}

func (s *Service) getCode(chain, address string) ([]byte, error) {
	if res, err := s.node.GetCodeAt(chain, address); err != nil {
		log.Println("Failed To Get Code", "chain", chain, "address", address)
		return nil, err
	} else {
		return res, nil
	}
}
