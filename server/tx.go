package server

import (
	gRPC "github.com/04Akaps/tx-sender-server/gRPC/client"
	"github.com/04Akaps/tx-sender-server/service"
	"github.com/04Akaps/tx-sender-server/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type tx struct {
	gRPCClient *gRPC.AuthGrpcClient
	service    *service.Service
}

func newTx(server *Server, gRPCClient *gRPC.AuthGrpcClient, service *service.Service) {
	t := &tx{
		gRPCClient: gRPCClient,
		service:    service,
	}

	baseURI := "/tx"

	server.register(baseURI+"/login", GET, t.login)

	server.register(baseURI+"/latest-block-number", GET, t.latestBlock)
	server.register(baseURI+"/block", GET, t.blockByNumber)
	server.register(baseURI+"/tx-receipt", GET, t.txReceipt)
	server.register(baseURI+"/balance", GET, verifyLogin(gRPCClient), t.balance)

	server.register(baseURI+"/unSign", POST, verifyLogin(gRPCClient), t.unSign)
	server.register(baseURI+"/sendTx", POST, verifyLogin(gRPCClient), t.sendTx)
}

func (t *tx) login(c *gin.Context) {
	var req types.LoginReq

	if err := c.ShouldBindQuery(&req); err != nil {
		response(c, http.StatusUnprocessableEntity, nil, err.Error())
	} else if res, err := t.gRPCClient.CreateAuth(req.Address); err != nil {
		response(c, http.StatusInternalServerError, nil, err.Error())
	} else {
		response(c, http.StatusOK, res, "")
	}
}

func (t *tx) blockByNumber(c *gin.Context) {
	var req types.BlockByNumberReq

	if err := c.ShouldBindQuery(&req); err != nil {
		response(c, http.StatusUnprocessableEntity, nil, err.Error())
	} else if res, err := t.service.GetBlockByNumber(req.Chain, req.Number); err != nil {
		response(c, http.StatusInternalServerError, nil, err.Error())
	} else {
		response(c, http.StatusOK, res, "")
	}
}

func (t *tx) txReceipt(c *gin.Context) {
	var req types.TxReceiptReq

	if err := c.ShouldBindQuery(&req); err != nil {
		response(c, http.StatusUnprocessableEntity, nil, err.Error())
	} else if res, err := t.service.GetTxReceipt(req.Chain, req.Tx); err != nil {
		response(c, http.StatusInternalServerError, nil, err.Error())
	} else {
		response(c, http.StatusOK, res, "")
	}
}

func (t *tx) balance(c *gin.Context) {
	var req types.ChainReq

	if err := c.ShouldBindQuery(&req); err != nil {
		response(c, http.StatusUnprocessableEntity, nil, err.Error())
	} else if from, err := getAddressByToken(c, t.gRPCClient); err != nil {
		response(c, http.StatusUnprocessableEntity, nil, err.Error())
	} else if res, err := t.service.GetBalance(req.Chain, from); err != nil {
		response(c, http.StatusInternalServerError, nil, err.Error())
	} else {
		response(c, http.StatusOK, res, "")
	}
}

func (t *tx) latestBlock(c *gin.Context) {
	var req types.ChainReq

	if err := c.ShouldBindQuery(&req); err != nil {
		response(c, http.StatusUnprocessableEntity, nil, err.Error())
	} else if res, err := t.service.GetLatestBlockNumber(req.Chain); err != nil {
		response(c, http.StatusInternalServerError, nil, err.Error())
	} else {
		response(c, http.StatusOK, res, "")
	}
}

// TODO -----

func (t *tx) unSign(c *gin.Context) {
	var req types.UnSignReq

	if err := c.ShouldBindJSON(&req); err != nil {
		response(c, http.StatusUnprocessableEntity, nil, err.Error())
	} else if from, err := getAddressByToken(c, t.gRPCClient); err != nil {
		response(c, http.StatusUnprocessableEntity, nil, err.Error())
	} else if res, err := t.service.UnSignTx(req, from); err != nil {
		response(c, http.StatusInternalServerError, nil, err.Error())
	} else {
		response(c, http.StatusOK, res, "")
	}

}

func (t *tx) sendTx(c *gin.Context) {}
