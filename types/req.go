package types

import "github.com/ethereum/go-ethereum/common/hexutil"

type LoginReq struct {
	Address string `form:"address" binding:"required"`
}

type ChainReq struct {
	Chain string `form:"chain" binding:"required"`
}

type BlockByNumberReq struct {
	Chain  string `form:"chain" binding:"required"`
	Number int64  `form:"number" binding:"required"`
}

type TxReceiptReq struct {
	Chain string `form:"chain" binding:"required"`
	Tx    string `form:"tx" binding:"required"`
}

type UnSignReq struct {
	To    string `json:"to" binding:"required"`
	Chain string `json:"chain" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type SignedTxReq struct {
	Chain string        `json:"chain" binding:"required"`
	Hash  string        `json:"hash" binding:"required"`
	Sign  hexutil.Bytes `json:"sign" binding:"required"`
	Tx    string        `json:"tx" bson:"tx"`
}
