package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/04Akaps/tx-sender-server/repository/node"
	"github.com/04Akaps/tx-sender-server/repository/redis"
	. "github.com/04Akaps/tx-sender-server/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"
	"log"
	"math/big"
	"strconv"
)

type Service struct {
	node  node.NodeImpl
	redis redis.RedisImpl
}

func NewService(node node.NodeImpl, redis redis.RedisImpl) *Service {
	return &Service{node, redis}
}

func (s *Service) UnSignTx(req UnSignReq, from string) (common.Hash, error) {
	log.Println("Make Hash For Send Value To EOA")

	if fromBalance, err := s.GetBalance(req.Chain, from); err != nil {
		return common.Hash{}, err
	} else if valueDecimal, err := decimal.NewFromString(req.Value); err != nil {
		log.Println("Failed to make value decimal")
		return common.Hash{}, err
	} else {
		decimalBalance := decimal.NewFromInt(fromBalance)

		if decimalBalance.Cmp(valueDecimal) < 0 {
			return common.Hash{}, errors.New(VerifyTokenErrMap[EnoughBalance])
		} else if valueDecimal.Cmp(decimal.Zero) == 0 {
			return common.Hash{}, errors.New(VerifyTokenErrMap[ZeroValueTransfer])
		} else if signer, err := s.node.GetSigner(req.Chain); err != nil {
			log.Println("Failed To Get Signer", err.Error())
			return common.Hash{}, err
		} else if nonce, err := s.node.GetNonce(req.Chain, from); err != nil {
			log.Println("Failed To Get Nonce", err.Error())
			return common.Hash{}, err
		} else if perGas, err := s.node.GetFeePerGas(req.Chain); err != nil {
			log.Println("Failed To Get Fee Per Gas", err.Error())
			// get Fee
			return common.Hash{}, err
		} else if baseFee, err := s.node.GetBaseFee(req.Chain); err != nil {
			log.Println("Failed To Get BaseFee", err.Error())
			return common.Hash{}, err
		} else {
			valueInt, _ := strconv.Atoi(req.Value)
			to := common.HexToAddress(req.To)

			dynamicTx := &etypes.DynamicFeeTx{
				To:        &to,
				Nonce:     nonce,
				Value:     big.NewInt(int64(valueInt)),
				GasFeeCap: perGas.Add(perGas, baseFee),
				GasTipCap: perGas,
				Data:      nil,
				V:         common.Big1,
				R:         common.Big1,
				S:         common.Big1,
			}

			if dynamicTx.Gas, err = s.node.GetDefaultGasLimit(req.Chain); err != nil {
				log.Println("Failed To Get DefaultGasLimit", err.Error())
				return common.Hash{}, err
			} else if dynamicTx.ChainID, err = s.node.GetChainID(req.Chain); err != nil {
				log.Println("Failed To Get ChainId", err.Error())
				return common.Hash{}, err
			} else {
				tx := etypes.NewTx(dynamicTx)
				return signer.Hash(tx), nil
			}
		}
	}

}

func (s *Service) SendSignedTx(req SignedTxReq, from string) error {
	hash := common.HexToHash(req.Hash)
	fromAddress := common.HexToAddress(from)

	if signer, err := s.node.GetSigner(req.Chain); err != nil {
		log.Println("Failed To Get Signer", err.Error())
		return err
	} else if publicKey, err := crypto.Ecrecover(hash[:], req.Sign); err != nil {
		log.Println("Failed To Ecrecover", err.Error())
		return err
	} else if txBytes, err := hexutil.Decode(req.Tx); err != nil {
		log.Println("Failed To Decode Tx", err.Error())
		return err
	} else {

		var unSignTx *etypes.Transaction
		recoveredPublicKey := crypto.Keccak256(publicKey[1:])[12:]

		if ok := bytes.Equal(fromAddress[:], recoveredPublicKey); !ok {
			return fmt.Errorf("Failed To From Valid")
		} else if err := json.Unmarshal(txBytes, &unSignTx); err != nil {
			log.Println("Failed To UnMarshal txBytes to unsignTx", err.Error())
			return err
		}

		r, _s, v, err := signer.SignatureValues(unSignTx, req.Sign)
		if err != nil {
			log.Println("Failed To SignatureValues", err.Error())
			return err
		}

		tx := etypes.NewTx(&etypes.DynamicFeeTx{
			ChainID:   unSignTx.ChainId(),
			Nonce:     unSignTx.Nonce(),
			GasTipCap: unSignTx.GasTipCap(),
			GasFeeCap: unSignTx.GasFeeCap(),
			Gas:       unSignTx.Gas(),
			To:        unSignTx.To(),
			Data:      unSignTx.Data(),
			Value:     unSignTx.Value(),
			V:         v,
			R:         r,
			S:         _s,
		})

		if err = s.node.SendTransaction(req.Chain, tx); err != nil {
			log.Println("Failed To Send Tx", err.Error())
			return err
		} else {
			return nil
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
