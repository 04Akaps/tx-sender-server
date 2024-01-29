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
	From    string `json:"from"`
	Address string `json:"address"`
	Chain   string `json:"chain"`
}
