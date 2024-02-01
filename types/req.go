package types

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
	To     string `json:"to"`
	Chain  string `json:"chain" binding:"required"`
	ABI    string `json:"abi"`
	Method string `json:"method"`
	Args   string `json:"args"`
	Value  string `json:"value"`
}
